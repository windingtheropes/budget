package main
import (
	"github.com/gin-gonic/gin"
	"github.com/windingtheropes/budget/auth"
	"github.com/windingtheropes/budget/argent"
	"github.com/windingtheropes/budget/dotenv"
	"github.com/windingtheropes/budget/based"
)

func main() {
	dotenv.Init()
	based.InitDB()
	
	// db := based.DB();
	engine := gin.Default();

	auth.LoadRoutes(engine);
	argent.LoadRoutes(engine);

	engine.Run("localhost:3000");
}