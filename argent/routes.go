package argent

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/windingtheropes/budget/auth"
	"github.com/windingtheropes/budget/json"
	"github.com/windingtheropes/budget/tables"
	"github.com/windingtheropes/budget/types"
)

// Authentication routes
func LoadRoutes(engine *gin.Engine) {
	// Get an entry and automatically hydrate it with tags
	engine.GET("/api/argent/transaction", func(ctx *gin.Context) {
		code, usr := auth.GetUserFromRequest(auth.GetTokenFromRequest(ctx))
		if code >= 400 {
			json.AbortWithStatusMessage(ctx, code, "")
			return
		}

		transactions, err := tables.Transaction.Get("user_id=?", usr.Id)
		if err != nil {
			fmt.Println(err)
			json.AbortWithStatusMessage(ctx, 500, "Internal error.")
			return
		}
		hydratedTransactions, err := HydrateTransactions(transactions)
		if err != nil {
			fmt.Println(err)
			json.AbortWithStatusMessage(ctx, 500, "Internal error.")
			return
		}
		ctx.AbortWithStatusJSON(200, json.ValueResponse[[]types.HydTransactionEntry]{Value: hydratedTransactions})
	})
	engine.POST("/api/argent/transaction/new", func(ctx *gin.Context) {
		code, usr := auth.GetUserFromRequest(auth.GetTokenFromRequest(ctx))
		if code >= 400 {
			json.AbortWithStatusMessage(ctx, code, "")
			return
		}

		var body json.TransactionForm
		if err := ctx.ShouldBindJSON(&body); err != nil {
			fmt.Printf("%v\n", err)
			json.AbortWithStatusMessage(ctx, 400, "Invalid JSON.")
			return
		}
		if body.Amount == 0 {
			json.AbortWithStatusMessage(ctx, 400, "Amount cannot be zero.")
		}
		transaction_id, err := tables.Transaction.New(types.TransactionEntryForm{
			User_Id:        usr.Id,
			Type_Id:        body.Type_Id,
			Amount:         body.Amount,
			Currency:       body.Currency,
			Msg:            body.Msg,
			Unix_Timestamp: body.Unix_Timestamp,
			Vendor:         body.Vendor,
		})
		if err != nil {
			fmt.Printf("%v\n", err)
			json.AbortWithStatusMessage(ctx, 500, "Internal Error.")
			return
		}

		// add tags
		for i := range body.Tags {
			tag_id := body.Tags[i]
			_, err := tables.TagAssignment.New(types.TagAssignmentForm{
				Tag_Id:         tag_id,
				Transaction_Id: transaction_id,
			})
			if err != nil {
				json.AbortWithStatusMessage(ctx, 500, "Internal Error")
				return
			}
		}

		// create budget entries
		for i := range body.Budget_Entries {
			budget_entry := body.Budget_Entries[i]
			if !BudgetExists(budget_entry.Budget_Id) {
				json.AbortWithStatusMessage(ctx, 400, "Transaction does not exist.")
				return
			}
			if !UserOwnsBudget(usr.Id, budget_entry.Budget_Id) {
				json.AbortWithStatusMessage(ctx, 403, "Budget not owned by user requesting it.")
				return
			}

			_, err := tables.BudgetEntry.New(types.BudgetEntryForm{
				Transaction_Id: transaction_id,
				Budget_Id:      budget_entry.Budget_Id,
				Amount:         budget_entry.Amount,
			})
			if err != nil {
				json.AbortWithStatusMessage(ctx, 500, "Internal Error")
				return
			}
		}

		json.AbortWithStatusMessage(ctx, 200, fmt.Sprintf("Entry added with ID %v", transaction_id))
	})
	engine.DELETE("/api/argent/transaction/delete", func(ctx *gin.Context) {
		code, usr := auth.GetUserFromRequest(auth.GetTokenFromRequest(ctx))
		if code >= 400 {
			json.AbortWithStatusMessage(ctx, code, "")
			return
		}

		var transaction_query = ctx.Request.URL.Query().Get("id")
		transaction_id, err := strconv.ParseInt(transaction_query, 0, 64)
		if err != nil {
			fmt.Printf("%v\n", err)
			json.AbortWithStatusMessage(ctx, 400, "Transaction ID must be an integer.")
			return
		}

		if !TransactionExists(transaction_id) {
			json.AbortWithStatusMessage(ctx, 400, "Transaction does not exist.")
			return
		}
		if !UserOwnsTransaction(usr.Id, transaction_id) {
			json.AbortWithStatusMessage(ctx, 403, "Access Denied.")
			return
		}

		tags, err := tables.TagAssignment.Get("entry_id = ?", transaction_id)
		if err != nil {
			fmt.Printf("%v\n", err)
			json.AbortWithStatusMessage(ctx, 500, "Internal Error.")
			return
		}

		budget_entries, err := tables.BudgetEntry.Get("transaction_id=?", transaction_id)
		if err != nil {
			json.AbortWithStatusMessage(ctx, 500, "Internal Error.")
			return
		}
		// Remove all tag assignments, because they are dependent on the existance of this transaction
		for i := range tags {
			tag := tags[i]
			if _, err := tables.TagAssignment.Delete("(tag_id=? AND entry_id=?)", tag.Tag_Id, transaction_id); err != nil {
				json.AbortWithStatusMessage(ctx, 500, "Internal Error.")
				return
			}
		}
		// Remove all budget entry assignments, because they are dependent on the existance of this transaction
		for i := range budget_entries {
			budget_entry := budget_entries[i]
			if _, err := tables.BudgetEntry.Delete("id=?", budget_entry.Id); err != nil {
				json.AbortWithStatusMessage(ctx, 500, "Internal Error.")
				return
			}
		}

		if _, err := tables.Transaction.Delete("id=?", transaction_id); err != nil {
			fmt.Printf("%v\n", err)
			json.AbortWithStatusMessage(ctx, 500, "Internal Error.")
			return
		}

		json.AbortWithStatusMessage(ctx, 200, fmt.Sprintf("Entry %v was deleted.", transaction_id))
	})
	// List user tags
	engine.GET("/api/argent/tag", func(ctx *gin.Context) {
		code, usr := auth.GetUserFromRequest(auth.GetTokenFromRequest(ctx))
		if code >= 400 {
			json.AbortWithStatusMessage(ctx, code, "")
			return
		}

		tags, err := GetUserTags(usr.Id)
		if err != nil {
			json.AbortWithStatusMessage(ctx, 500, "Internal error.")
			return
		}
		ctx.AbortWithStatusJSON(200, json.ValueResponse[[]types.Tag]{Value: tags})
	})
	// List user budgets
	engine.GET("/api/argent/budget", func(ctx *gin.Context) {
		code, usr := auth.GetUserFromRequest(auth.GetTokenFromRequest(ctx))
		if code >= 400 {
			json.AbortWithStatusMessage(ctx, code, "")
			return
		}

		budgets, err := tables.Budget.Get("user_id=?", usr.Id)
		if err != nil {
			json.AbortWithStatusMessage(ctx, 500, "Internal error.")
			return
		}
		hydBudgets, err := HydrateBudgetsWithTagBudgets(budgets)
		if err != nil {
			json.AbortWithStatusMessage(ctx, 500, "Internal error.")
			return
		}

		ctx.AbortWithStatusJSON(200, json.ValueResponse[[]types.HydBudget]{Value: hydBudgets})
	})
	engine.POST("/api/argent/tag/new", func(ctx *gin.Context) {
		code, usr := auth.GetUserFromRequest(auth.GetTokenFromRequest(ctx))
		if code >= 400 {
			json.AbortWithStatusMessage(ctx, code, "")
			return
		}

		var body json.TagForm
		if err := ctx.ShouldBindJSON(&body); err != nil {
			json.AbortWithStatusMessage(ctx, 400, "Invalid JSON.")
			return
		}
		if UserTagNameExists(body.Name, usr.Id) {
			json.AbortWithStatusMessage(ctx, 400, "Tag exists.")
			return
		}

		tag_id, err := NewUserTag(body.Name, usr.Id)
		if err != nil {
			json.AbortWithStatusMessage(ctx, 500, "Internal error.")
			return
		}

		json.AbortWithStatusMessage(ctx, 200, fmt.Sprintf("Created tag %v (%v).", body.Name, tag_id))
	})
	engine.POST("/api/argent/budget/new", func(ctx *gin.Context) {
		code, usr := auth.GetUserFromRequest(auth.GetTokenFromRequest(ctx))
		if code >= 400 {
			json.AbortWithStatusMessage(ctx, code, "")
			return
		}

		var body json.BudgetForm

		if err := ctx.ShouldBindJSON(&body); err != nil {
			json.AbortWithStatusMessage(ctx, 400, "Invalid JSON.")
			return
		}
		if UserBudgetNameExists(body.Name, usr.Id) {
			json.AbortWithStatusMessage(ctx, 400, "Budget exists.")
			return
		}
		budget_id, err := tables.Budget.New(types.BudgetForm{
			Name:    body.Name,
			User_Id: usr.Id,
			Goal:    *body.Goal,
		})
		if err != nil {
			json.AbortWithStatusMessage(ctx, 500, "Internal error.")
			return
		}
		// Try to load each tagbudget
		for i := range body.Tag_Budgets {
			tagBudget := body.Tag_Budgets[i]
			if !TagExists(tagBudget.Tag_Id) {
				json.AbortWithStatusMessage(ctx, 400, "Tag does not exist.")
				return
			}
			if !UserOwnsTag(usr.Id, tagBudget.Tag_Id) {
				json.AbortWithStatusMessage(ctx, 403, "Access Denied.")
				return
			}

			// create tagbudget
			_, err := tables.TagBudget.New(types.TagBudgetForm{
				Tag_Id:    tagBudget.Tag_Id,
				Budget_Id: budget_id,
				Goal:      tagBudget.Goal,
				Type_Id:   tagBudget.Type_Id,
			})
			if err != nil {
				json.AbortWithStatusMessage(ctx, 500, "Internal Error")
				return
			}
		}
		json.AbortWithStatusMessage(ctx, 200, fmt.Sprintf("Created budget %v (%v).", body.Name, budget_id))
	})
	// THESE NEED TO BE CACHED
	engine.GET("/api/argent/currency/exchange", func(ctx *gin.Context) {
		query := ctx.Query("currency")
		if query == "" {
			json.AbortWithStatusMessage(ctx, 400, "Malformatted request.")
			return
		}

		rate, err := GetExchangeCAD(query)
		if err != nil {
			json.AbortWithStatusMessage(ctx, 500, "Internal Error.")
			return
		}
		ctx.AbortWithStatusJSON(200, json.ValueResponse[string]{Value: fmt.Sprintf("%v", rate)})
	})
	engine.GET("/api/argent/currency", func(ctx *gin.Context) {
		currencies, err := GetCurrencies()
		if err != nil {
			json.AbortWithStatusMessage(ctx, 500, "Internal Error.")
			return
		}
		ctx.AbortWithStatusJSON(200, json.ValueResponse[[]string]{Value: currencies})
	})
	engine.GET("/api/argent/type", func(ctx *gin.Context) {
		transactionTypes, err := tables.TransactionType.Get("*")
		if err != nil {
			json.AbortWithStatusMessage(ctx, 500, "Internal Error.")
			return
		}
		ctx.AbortWithStatusJSON(200, json.ValueResponse[[]types.TransactionType]{Value: transactionTypes})
	})
}
