package auth

// Wrappers reduce boilerplate code by handing errors and context aborts from the gin context in-function

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/windingtheropes/budget/json"
	"github.com/windingtheropes/budget/tables"
	"github.com/windingtheropes/budget/types"
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
func GetUserFromRequest(token string) (int, types.User) {
	var usr types.User;
	s, err := tables.Session.Get("token=?", token)
	if err != nil {
		return 500, usr
	}
	if len(s) == 0 {
		// No session exists
		return 403, usr
	}
	session := s[0]
	if !IsValidSession(&session) {
		// Token expired
		return 403, usr
	}
	usrs, err := tables.User.Get("id=?", session.User_Id)
	if err != nil {
		return 500, usr
	}
	if len(usrs) == 0 {
		// User doesn't exist
		return 403, usr
	}
	return 200, usr
}
