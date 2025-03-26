package json
import (
	"github.com/gin-gonic/gin"
)

type GenericResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
type SessionResponse struct {
	Code    int    `json:"code"`
	Token   string `json:"token"`
}
type ListResponse struct {
	Value []string `json:"value"`
}
type ValueResponse struct {
	Value string `json:"value"`
}

func NewResponse(ctx *gin.Context, code int, message string) {
	ctx.JSON(code, GenericResponse{
		Code:    code,
		Message: message,
	})
	ctx.Abort();
}

