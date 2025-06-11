package auth

// Wrappers reduce boilerplate code by handing errors and context aborts from the gin context in-function

import (
	"strings"
	"github.com/gin-gonic/gin"
	"github.com/windingtheropes/budget/types"
	"github.com/windingtheropes/budget/json"
)

// Parse the token from the http authorization header
func GetTokenFromRequest(ctx *gin.Context) string {
	a := ctx.Request.Header.Get("Authorization")
	authorization := strings.Split(a, " ")
	if authorization[0] != "Bearer" || len(authorization) < 2 {
		json.AbortWithStatusMessage(ctx, 400, "Invalid authorization header.")
	}
	return authorization[1]
} 

// Authentication middleware, returns either ([200], [user]), ([4-5xx], nil)
func GetUserFromRequestNew(token string) (int, []types.User) {
	s, err := SessionTable.Get("token=?", token)
	if err != nil {
		return 500, nil
	}
	if len(s) == 0 {
		// No session exists
		return 403, nil
	}  
	session := s[0]
	if !IsValidSession(&session) {
		// Token expired
		return 403, nil
	}
	usrs, err := UserTable.Get("id=?", session.User_Id)
	if err != nil {
		return 500, nil
	}
	if len(usrs) == 0 {
		// User doesn't exist
		return 403, nil
	}
	return 200, usrs
}