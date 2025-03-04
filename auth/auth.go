package auth

import (
	"github.com/gin-gonic/gin"
	// "crypto"
)

type AuthForm struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}
type LoginForm struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}
type GenericResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
type LoginResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Token   string `json:"token"`
}

func LoadRoutes(engine *gin.Engine) {
	engine.POST("/account/new", func(ctx *gin.Context) {
		body := AuthForm{}
		if err := ctx.ShouldBindJSON(&body); err != nil {
			ctx.JSON(400, GenericResponse{
				Code:    400,
				Message: "Invalid JSON.",
			})
			ctx.Abort()
			return
		}

		// CHECK IF USER EXISTS
		// ctx.JSON(403, GenericResponse{
		// 	Code:    403,
		// 	Message: "User Exists.",
		// })

		ctx.JSON(200, GenericResponse{
			Code:    200,
			Message: "Created.",
		})
	})
	engine.GET("/account/login", func(ctx *gin.Context) {
		// CHECK CREDENTIALS
		// ctx.JSON(403, GenericResponse{
		// 	Code:    403,
		// 	Message: "Invalid Credentials.",
		// })

		// IF CORRECT
		// ctx.JSON(200, GenericResponse{
		// 	Code:    200,
		// 	Token: "NEWTOKENGEN",
		// })
		ctx.Status(403)
	})
}
