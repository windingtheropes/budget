package auth

import (
	"crypto/rand"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"

	b64 "encoding/base64"
	"log"

	"github.com/windingtheropes/budget/json"
)

// Authentication routes
func LoadRoutes(engine *gin.Engine) {
	// New Account
	engine.POST("/api/account/new", func(ctx *gin.Context) {
		body := json.NewAccountForm{}
		if err := ctx.ShouldBindJSON(&body); err != nil {
			json.NewResponse(ctx, 400, "Invalid JSON.");
			return
		}
		
		users, err := GetUser(body.Email)
		if err != nil {
			log.Fatal(err)
			json.NewResponse(ctx, 500, "Interal error.");
			return
		}
		if len(users) != 0 {
			json.NewResponse(ctx, 403, "Email already in use.");
			return
		}

		// Password not hashed
		user_id, err := AddUser(body.Name, body.Email, body.Password);
		if err != nil {
			log.Fatal(err)
			json.NewResponse(ctx, 500, "Interal error.");
			return
		}

		json.NewResponse(ctx, 200, fmt.Sprintf("Created user %d", user_id));
	})

	// Login gives a session token
	engine.PUT("/api/account/login", func(ctx *gin.Context) {
		body := json.LoginForm{}
		// Bind the json to the loginform body, or return an error
		if err := ctx.ShouldBindJSON(&body); err != nil {
			json.NewResponse(ctx, 400, "Invalid JSON.");
			return
		}
		
		// catch unknown errors
		users, err := GetUser(body.Email)
		if err != nil {
			log.Fatal(err)
			json.NewResponse(ctx, 500, "Interal error.");
			return
		}
		if len(users) == 1 {
			usr := users[0];
			if body.Password == usr.Password {
				t, _, err := NewSession(usr.Id)
				var token string = t[0]

				if err != nil {
					log.Fatal(err)
					json.NewResponse(ctx, 500, "Interal error.");
					return
				}

				ctx.AbortWithStatusJSON(200, json.SessionResponse{
					Code: 200,
					Token: token,
				})
				return
			} else {
				// password incorrect
				json.NewResponse(ctx, 403, "Invalid credentials.");
				return
			}
		} else if len(users) == 0 {
			// user doesn't exist with email
			json.NewResponse(ctx, 403, "Invalid credentials.");
			return
		}
	})
}

// Returns whether the token belongs to a session, and if so, the user id which it belongs to.
func IsSession(token string) (bool, int) {
	sessions, err := GetSession(token);
	if err != nil {
		return false, 0
	}
	if len(sessions) == 0 {
		return false, 0
	}
	
	session := sessions[0];
	return true, session.User_Id
}
func GetAuthorization(ctx *gin.Context) int {
	a := ctx.Request.Header.Get("Authorization");
	authorization := strings.Split(a, " ")
	if authorization[0] != "Bearer" || len(authorization) != 2 {
		json.NewResponse(ctx, 400, "Invalid authorization header.")
	}
	token := authorization[1]

	isValid, user_id := IsSession(token); 
	if !isValid {
		json.NewResponse(ctx, 403, "Not allowed.")
	} 
	return user_id
}

func GenToken(length int) string {
	var keyBytes = make([]byte, length);
	rand.Read(keyBytes);
	return b64.RawStdEncoding.EncodeToString(keyBytes);
}
