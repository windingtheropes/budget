package auth

// Wrappers reduce boilerplate code by handing errors and context aborts from the gin context in-function

import (
	"strings"
	"github.com/gin-gonic/gin"
	"github.com/windingtheropes/budget/types"
	"github.com/windingtheropes/budget/json"
)

// Parse the token from the http authorization header
func getTokenFromRequest(ctx *gin.Context) string {
	a := ctx.Request.Header.Get("Authorization")
	authorization := strings.Split(a, " ")
	if authorization[0] != "Bearer" || len(authorization) < 2 {
		json.AbortWithStatusMessage(ctx, 400, "Invalid authorization header.")
	}
	return authorization[1]
} 
// Get a session from an http request using authorization
func getSessionFromRequest(ctx *gin.Context) types.Session {
	s, err := GetSession(getTokenFromRequest(ctx))
	if err != nil {
		json.AbortWithStatusMessage(ctx, 500, "Internal Error.")
	}
	if len(s) == 0 {
		// No session exists
		json.AbortWithStatusMessage(ctx, 403, "Not allowed.")
	}  
	session := s[0]
	if !session.IsValid() {
		// Token expired
		json.AbortWithStatusMessage(ctx, 403, "Not allowed.")
	}
	return session
}
// Full authentication pipeline, return a user
func GetUserFromRequest(ctx *gin.Context) types.User {
	session := getSessionFromRequest(ctx);
	usrs, err := GetUser(types.UserID(session.User_Id));
	if err != nil {
		json.AbortWithStatusMessage(ctx, 500, "Internal error.")
	}
	if len(usrs) == 0 {
		json.AbortWithStatusMessage(ctx, 403, "Invalid credentials.")
	}
	return usrs[0]
}