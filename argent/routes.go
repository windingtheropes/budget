package argent

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/windingtheropes/budget/auth"
	"github.com/windingtheropes/budget/json"
	"github.com/windingtheropes/budget/types"
)

// Authentication routes
func LoadRoutes(engine *gin.Engine) {
	// Get an entry and automatically hydrate it with tags
	engine.GET("/api/argent/entry", func(ctx *gin.Context) {
		code, usrs := auth.GetUserFromRequestNew(auth.GetTokenFromRequest(ctx))
		if code >= 400 {
			json.AbortWithStatusMessage(ctx, code, "")
			return
		}
		usr := usrs[0]

		transactions, err := GetUserTransactions(usr.Id)
		if err != nil {
			fmt.Println(err)
			json.AbortWithStatusMessage(ctx, 500, "Internal error.")
			return
		}
		hydratedTransactions, err := HydrateTransactionsWithTags(transactions)
		if err != nil {
			fmt.Println(err)
			json.AbortWithStatusMessage(ctx, 500, "Internal error.")
			return
		}
		ctx.AbortWithStatusJSON(200, json.ValueResponse[[]types.HydTransactionEntry]{Value: hydratedTransactions})
	})
	engine.POST("/api/argent/entry/new", func(ctx *gin.Context) {
		code, usrs := auth.GetUserFromRequestNew(auth.GetTokenFromRequest(ctx))
		if code >= 400 {
			json.AbortWithStatusMessage(ctx, code, "")
			return
		}
		usr := usrs[0]

		var body json.NewTransactionForm
		if err := ctx.ShouldBindJSON(&body); err != nil {
			fmt.Printf("%v\n", err)
			json.AbortWithStatusMessage(ctx, 400, "Invalid JSON.")
			return
		}
		if body.Amount == 0 {
			json.AbortWithStatusMessage(ctx, 400, "Amount cannot be zero.")
		}
		id, err := NewTransaction(usr.Id, body.Type_Id, body.Amount, body.Currency, body.Msg, body.Unix_Timestamp, body.Vendor)
		if err != nil {
			fmt.Printf("%v\n", err)
			json.AbortWithStatusMessage(ctx, 500, "Internal Error.")
			return
		}

		if len(body.Tags) > 0 {
			if err := AddTagsById(id, body.Tags); err != nil {
				json.AbortWithStatusMessage(ctx, 500, "Internal Error.")
				return
			}
		}

		json.AbortWithStatusMessage(ctx, 200, fmt.Sprintf("Entry added with ID %v", id))
	})
	engine.DELETE("/api/argent/entry/delete", func(ctx *gin.Context) {
		code, usrs := auth.GetUserFromRequestNew(auth.GetTokenFromRequest(ctx))
		if code >= 400 {
			json.AbortWithStatusMessage(ctx, code, "")
			return
		}
		usr := usrs[0]

		// TODO relatively unsafe
		var transaction_query = ctx.Request.URL.Query().Get("id")
		_tid, err := strconv.ParseInt(transaction_query, 0, 64)
		if err != nil {
			fmt.Printf("%v\n", err)
			json.AbortWithStatusMessage(ctx, 400, "Transaction ID must be an integer.")
			return
		}
		var transaction_id = int(_tid)

		if !TransactionExists(transaction_id) {
			json.AbortWithStatusMessage(ctx, 400, "Transaction does not exist.")
			return
		}
		if !UserOwnsTransaction(usr.Id, transaction_id) {
			json.AbortWithStatusMessage(ctx, 403, "Access Denied.")
			return
		}

		tags, err := GetTagAssignments(transaction_id)
		if err != nil {
			fmt.Printf("%v\n", err)
			json.AbortWithStatusMessage(ctx, 500, "Internal Error.")
			return
		}
		// Remove all tag assignments, because they are dependent on the existance of this entry
		for i := 0; i < len(tags); i++ {
			tag := tags[i]
			if _, err := DeleteTagOnEntry(tag.Id, transaction_id); err != nil {
				fmt.Printf("%v\n", err)
				json.AbortWithStatusMessage(ctx, 500, "Internal Error.")
				return
			}
		}

		if _, err := DeleteTransaction(transaction_id); err != nil {
			fmt.Printf("%v\n", err)
			json.AbortWithStatusMessage(ctx, 500, "Internal Error.")
			return
		}

		json.AbortWithStatusMessage(ctx, 200, fmt.Sprintf("Entry %v was deleted.", transaction_id))
	})
	// List user tags
	engine.GET("/api/argent/tag", func(ctx *gin.Context) {
		code, usrs := auth.GetUserFromRequestNew(auth.GetTokenFromRequest(ctx))
		if code >= 400 {
			json.AbortWithStatusMessage(ctx, code, "")
			return
		}
		usr := usrs[0]

		tags, err := GetUserTags(usr.Id)
		if err != nil {
			json.AbortWithStatusMessage(ctx, 500, "Internal error.")
			return
		}
		hydTags, err := HydrateTagsWithTagBudgets(tags)
		if err != nil {
			json.AbortWithStatusMessage(ctx, 500, "Internal error.")
			return
		}
		ctx.AbortWithStatusJSON(200, json.ValueResponse[[]types.HydTag]{Value: hydTags})
	})
	// List user budgets
	engine.GET("/api/argent/budget", func(ctx *gin.Context) {
		code, usrs := auth.GetUserFromRequestNew(auth.GetTokenFromRequest(ctx))
		if code >= 400 {
			json.AbortWithStatusMessage(ctx, code, "")
			return
		}
		usr := usrs[0]

		budgets, err := GetUserBudgets(usr.Id)
		if err != nil {
			json.AbortWithStatusMessage(ctx, 500, "Internal error.")
			return
		}
		if err != nil {
			json.AbortWithStatusMessage(ctx, 500, "Internal error.")
			return
		}
		ctx.AbortWithStatusJSON(200, json.ValueResponse[[]types.Budget]{Value: budgets})
	})
	engine.POST("/api/argent/tag/new", func(ctx *gin.Context) {
		code, usrs := auth.GetUserFromRequestNew(auth.GetTokenFromRequest(ctx))
		if code >= 400 {
			json.AbortWithStatusMessage(ctx, code, "")
			return
		}
		usr := usrs[0]

		var body json.NewTagForm
		if err := ctx.ShouldBindJSON(&body); err != nil {
			json.AbortWithStatusMessage(ctx, 400, "Invalid JSON.")
			return
		}
		if UserTagNameExists(body.Name, usr.Id) {
			json.AbortWithStatusMessage(ctx, 400, "Tag exists.")
			return
		}
		id, err := NewUserTag(body.Name, usr.Id)
		if err != nil {
			json.AbortWithStatusMessage(ctx, 500, "Internal error.")
			return
		}
		json.AbortWithStatusMessage(ctx, 200, fmt.Sprintf("Created tag %v (%v).", body.Name, id))
	})
	engine.POST("/api/argent/budget/new", func(ctx *gin.Context) {
		code, usrs := auth.GetUserFromRequestNew(auth.GetTokenFromRequest(ctx))
		if code >= 400 {
			json.AbortWithStatusMessage(ctx, code, "")
			return
		}
		usr := usrs[0]

		var body json.NewBudgetForm
		if err := ctx.ShouldBindJSON(&body); err != nil {
			json.AbortWithStatusMessage(ctx, 400, "Invalid JSON.")
			return
		}
		if UserBudgetNameExists(body.Name, usr.Id) {
			json.AbortWithStatusMessage(ctx, 400, "Budget exists.")
			return
		}
		id, err := NewBudget(body.Name, usr.Id, body.Type_Id, body.Goal)
		if err != nil {
			json.AbortWithStatusMessage(ctx, 500, "Internal error.")
			return
		}
		json.AbortWithStatusMessage(ctx, 200, fmt.Sprintf("Created budget %v (%v).", body.Name, id))
	})
	engine.POST("/api/argent/budget/entry/new", func(ctx *gin.Context) {
		code, usrs := auth.GetUserFromRequestNew(auth.GetTokenFromRequest(ctx))
		if code >= 400 {
			json.AbortWithStatusMessage(ctx, code, "")
			return
		}
		usr := usrs[0]

		var body json.NewBudgetEntryForm
		if err := ctx.ShouldBindJSON(&body); err != nil {
			json.AbortWithStatusMessage(ctx, 400, "Invalid JSON.")
			return
		}

		// Confirm ownership of requested resources.
		if !TransactionExists(body.Transaction_Id) {
			json.AbortWithStatusMessage(ctx, 400, "Transaction does not exist.")
			return
		}
		if !UserOwnsTransaction(usr.Id, body.Transaction_Id) {
			json.AbortWithStatusMessage(ctx, 403, "Access Denied.")
			return
		}
		if !BudgetExists(body.Transaction_Id) {
			json.AbortWithStatusMessage(ctx, 400, "Transaction does not exist.")
			return
		}
		if !UserOwnsBudget(usr.Id, body.Budget_Id) {
			json.AbortWithStatusMessage(ctx, 403, "Access Denied.")
			return
		}

		id, err := NewBudgetEntry(body.Transaction_Id, body.Budget_Id, body.Amount)
		if err != nil {
			json.AbortWithStatusMessage(ctx, 500, "Internal Error")
			return
		}

		json.AbortWithStatusMessage(ctx, 200, fmt.Sprintf("Created budget entry  with id %v", id))
	})
	// New TagBudget
	engine.POST("/api/argent/tag/budget/new", func(ctx *gin.Context) {
		code, usrs := auth.GetUserFromRequestNew(auth.GetTokenFromRequest(ctx))
		if code >= 400 {
			json.AbortWithStatusMessage(ctx, code, "")
			return
		}
		usr := usrs[0]

		var body json.NewTagBudgetForm
		if err := ctx.ShouldBindJSON(&body); err != nil {
			json.AbortWithStatusMessage(ctx, 400, "Invalid JSON.")
			return
		}

		if !TagExists(body.Tag_Id) {
			json.AbortWithStatusMessage(ctx, 400, "Tag does not exist.")
			return
		}
		if !UserOwnsTag(usr.Id, body.Tag_Id) {
			json.AbortWithStatusMessage(ctx, 403, "Access Denied.")
			return
		}
		if !BudgetExists(body.Budget_Id) {
			json.AbortWithStatusMessage(ctx, 400, "Budget does not exist.")
			return
		}
		if !UserOwnsBudget(usr.Id, body.Budget_Id) {
			json.AbortWithStatusMessage(ctx, 403, "Access Denied.")
			return
		}

		id, err := NewTagBudget(body.Tag_Id, body.Budget_Id, body.Goal, body.Type_Id)
		if err != nil {
			json.AbortWithStatusMessage(ctx, 500, "Internal Error")
			return
		}

		json.AbortWithStatusMessage(ctx, 200, fmt.Sprintf("Created budget entry  with id %v", id))
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
		ctx.AbortWithStatusJSON(200, json.ListResponse{Value: currencies})
	})
	engine.GET("/api/argent/type", func(ctx *gin.Context) {
		types, err := GetTransactionTypes()
		if err != nil {
			json.AbortWithStatusMessage(ctx, 500, "Internal Error.")
			return
		}
		ctx.AbortWithStatusJSON(200, json.TransactionTypesResponse{Value: types})
	})
}
