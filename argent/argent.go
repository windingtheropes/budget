package argent

import (
	"fmt"

	"github.com/gin-gonic/gin"
	// "github.com/windingtheropes/budget/auth"
	"github.com/windingtheropes/budget/json"
)

// Authentication routes
func LoadRoutes(engine *gin.Engine) {
	// New Account
	engine.GET("/argent/entries", func(ctx *gin.Context) {
		// uid := auth.GetAuthorization(ctx)
	})

	// THESE NEED TO BE CACHED
	engine.GET("/api/argent/currency/exchange", func(ctx *gin.Context) {
		query := ctx.Query("currency")
		if query == "" {
			json.NewResponse(ctx, 400, "Malformatted request.");
			return
		}

		rate, err := GetExchangeCAD(query);
		if err != nil {
			json.NewResponse(ctx, 500, "Internal Error.");
			return
		}
		ctx.AbortWithStatusJSON(200, json.ValueResponse{Value: fmt.Sprintf("%v", rate)})
	})
	engine.GET("/api/argent/currency", func(ctx *gin.Context) {
		currencies, err := GetCurrencies();
		if err != nil {
			json.NewResponse(ctx, 500, "Internal Error.");
			return
		}
		ctx.AbortWithStatusJSON(200, json.ListResponse{Value: currencies})
	})
}
