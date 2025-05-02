package main

import (
	"github.com/gin-gonic/gin"
	"github.com/windingtheropes/budget/argent"
	"github.com/windingtheropes/budget/auth"
	"github.com/windingtheropes/budget/based"
	"github.com/windingtheropes/budget/dotenv"
)

func main() {
	dotenv.Init()
	based.InitDB()

	// db := based.DB();
	engine := gin.Default()
	engine.Use(CORSMiddleware())
	auth.LoadRoutes(engine)
	argent.LoadRoutes(engine)

	engine.Run("localhost:3000")
}
// TEMP BYPASS ALL
func CORSMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
        c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
        c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
        c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(204)
            return
        }

        c.Next()
    }
}