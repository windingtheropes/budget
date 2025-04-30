package argent

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/windingtheropes/budget/auth"
	"github.com/windingtheropes/budget/json"
	"github.com/windingtheropes/budget/types"
)

// Authentication routes
func LoadRoutes(engine *gin.Engine) {
	// New Account
	engine.GET("/api/argent/entry", func(ctx *gin.Context) {
		code, usrs := auth.GetUserFromRequestNew(auth.GetTokenFromRequest(ctx))
		if code >= 400 {
			json.AbortWithStatusMessage(ctx, code, "")
			return
		}
		usr := usrs[0]

		transactions, err := GetTransactions(usr.Id)
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
		ctx.AbortWithStatusJSON(200, json.HydratedTransactionsResponse{Value: hydratedTransactions})
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
			fmt.Printf("%v\n", err);
			json.AbortWithStatusMessage(ctx, 400, "Invalid JSON.")
			return
		}
		if body.Amount == 0 {
			json.AbortWithStatusMessage(ctx, 400, "Amount cannot be zero.")
		}
		id, err := NewTransaction(usr.Id, body.Type_Id, body.Amount, body.Currency, body.Msg, body.Unix_Timestamp, body.Vendor)
		if err != nil {
			fmt.Printf("%v\n", err);
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

	engine.GET("/api/argent/tag", func(ctx *gin.Context) {
		code, usrs := auth.GetUserFromRequestNew(auth.GetTokenFromRequest(ctx))
		if code >= 400 {
			json.AbortWithStatusMessage(ctx, code, "")
			return
		}
		usr := usrs[0]

		tags, err := GetTag(types.UserID(usr.Id))
		if err != nil {
			json.AbortWithStatusMessage(ctx, 500, "Internal error.")
			return
		}
		ctx.AbortWithStatusJSON(200, json.TagResponse{Value: tags})
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
		if TagExists(body.Name, usr.Id) {
			json.AbortWithStatusMessage(ctx, 400, "Tag exists.")
			return
		}
		id, err := NewTag(usr.Id, body.Name)
		if err != nil {
			json.AbortWithStatusMessage(ctx, 500, "Internal error.")
			return
		}
		json.AbortWithStatusMessage(ctx, 200, fmt.Sprintf("Created tag %v (%v).", body.Name, id))
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
		ctx.AbortWithStatusJSON(200, json.ValueResponse{Value: fmt.Sprintf("%v", rate)})
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
