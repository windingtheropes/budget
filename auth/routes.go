package auth

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/windingtheropes/budget/json"
	"github.com/windingtheropes/budget/types"
)

// Authentication routes
func LoadRoutes(engine *gin.Engine) {
	// New Account
	engine.POST("/api/account/new", func(ctx *gin.Context) {
		body := json.NewAccountForm{}
		if err := ctx.ShouldBindJSON(&body); err != nil {
			json.AbortWithStatusMessage(ctx, 400, "Invalid JSON.")
			return
		}

		users, err := GetUser(types.Email(body.Email))
		if err != nil {
			log.Fatal(err)
			json.AbortWithStatusMessage(ctx, 500, "Interal error.")
			return
		}
		if len(users) != 0 {
			json.AbortWithStatusMessage(ctx, 403, "Email already in use.")
			return
		}

		// Password not hashed
		user_id, err := AddUser(body.Name, body.Email, body.Password)
		if err != nil {
			log.Fatal(err)
			json.AbortWithStatusMessage(ctx, 500, "Interal error.")
			return
		}

		json.AbortWithStatusMessage(ctx, 200, fmt.Sprintf("Created user %d", user_id))
	})

	// Login gives a session token
	engine.PUT("/api/account/login", func(ctx *gin.Context) {
		body := json.LoginForm{}
		// Bind the json to the loginform body, or return an error
		if err := ctx.ShouldBindJSON(&body); err != nil {
			json.AbortWithStatusMessage(ctx, 400, "Invalid JSON.")
			return
		}

		// catch unknown errors
		users, err := GetUser(types.Email(body.Email))
		if err != nil {
			log.Fatal(err)
			json.AbortWithStatusMessage(ctx, 500, "Interal error.")
			return
		}
		if len(users) == 1 {
			usr := users[0]
			if body.Password == usr.Password {
				expiry := (60 * 60 * 4)
				t, _, err := NewSession(usr.Id, expiry)
				if err != nil {
					log.Fatal(err)
					json.AbortWithStatusMessage(ctx, 500, "Interal error.")
					return
				}
				var token string = t[0]

				ctx.AbortWithStatusJSON(200, json.SessionResponse{
					Code:  200,
					Token: token,
				})
				return
			} else {
				// password incorrect
				json.AbortWithStatusMessage(ctx, 403, "Invalid credentials.")
				return
			}
		} else if len(users) == 0 {
			// user doesn't exist with email
			json.AbortWithStatusMessage(ctx, 403, "Invalid credentials.")
			return
		}
	})

	// Get info on a session from the enclosed token
	engine.PUT("/api/account/session", func(ctx *gin.Context) {
		GetUserFromRequest(ctx)
		json.AbortWithStatusMessage(ctx, 200, "Authorized.")
	})

	// Get user info
	engine.PUT("/api/account/user", func(ctx *gin.Context) {
		usr := GetUserFromRequest(ctx)

		ctx.AbortWithStatusJSON(200, json.UserInfoResponse{
			Id:    usr.Id,
			Name:  usr.Name,
			Email: usr.Email,
		})
	})
}
