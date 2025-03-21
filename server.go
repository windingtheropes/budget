package main
import (
	"github.com/gin-gonic/gin"
	"github.com/windingtheropes/budget/auth"
	"github.com/windingtheropes/budget/dotenv"
	"github.com/windingtheropes/budget/based"
	// "crypto"
)

func main() {
	dotenv.Init()
	based.InitDB()
	
	engine := gin.Default();
	auth.LoadRoutes(engine);
	
	engine.Run("localhost:3000");
}