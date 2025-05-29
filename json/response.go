package json

import (
	"github.com/gin-gonic/gin"
)

type GenericResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
type SessionResponse struct {
	Code  int    `json:"code"`
	Token string `json:"token"`
}
type ValueResponse[T any] struct {
	Value T `json:"value"`
}
type UserInfoResponse struct {
	Id    int64  `json:"id"`
	First_Name  string `json:"first_name"`
	Last_Name string `json:"last_name"`
	Email string `json:"email"`
}

func AbortWithStatusMessage(ctx *gin.Context, code int, message string) {
	ctx.AbortWithStatusJSON(code, GenericResponse{
		Code:    code,
		Message: message,
	})
}
