package auth

import (
	"time"

	"github.com/windingtheropes/budget/types"
)
func IsValidSession(session *types.Session) bool {
	return time.Now().Unix() <= session.Expiry
}