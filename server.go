package main
import (
	"github.com/gin-gonic/gin"
	"github.com/windingtheropes/budget/auth"
	// "crypto"
)

func main() {
	engine := gin.Default();
	auth.LoadRoutes(engine);
	engine.GET("/hello", func(ctx *gin.Context) {
		ctx.Status(200);
	})
	engine.Run("localhost:3000");
}