package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/windingtheropes/budget/argent"
	"github.com/windingtheropes/budget/auth"
	"github.com/windingtheropes/budget/based"
	"github.com/windingtheropes/budget/dotenv"
)

func main() {
	dotenv.Init()
	based.InitDB()
	engine := gin.Default()
	engine.Use(CORSMiddleware())
	auth.LoadRoutes(engine)
	argent.LoadRoutes(engine)
	engine.SetTrustedProxies(nil)
	// for docker deployment, specifying a different ip to listen on
	var ip string;
	if os.Getenv("WEBIP") != "" {
		ip = os.Getenv("WEBIP")
	} else {
		ip = "localhost"
	}

	engine.Run(fmt.Sprintf("%v:%v", ip ,os.Getenv("WEBPORT")))
}

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