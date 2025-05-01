package json

import (
	"github.com/gin-gonic/gin"
	"github.com/windingtheropes/budget/types"
)

type GenericResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
type SessionResponse struct {
	Code  int    `json:"code"`
	Token string `json:"token"`
}
type ListResponse struct {
	Value []string `json:"value"`
}
type TransactionTypesResponse struct {
	Value []types.TransactionType `json:"value"`
}
type ValueResponse struct {
	Value string `json:"value"`
}
type HydratedTransactionsResponse struct {
	Value []types.HydTransactionEntry `json:"value"`
}
type TagResponse struct {
	Value []types.Tag `json:"value"`
}
type UserInfoResponse struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
}

func AbortWithStatusMessage(ctx *gin.Context, code int, message string) {
	ctx.AbortWithStatusJSON(code, GenericResponse{
		Code:    code,
		Message: message,
	})
}
