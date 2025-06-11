package auth

import (
	"crypto/rand"
	b64 "encoding/base64"

	"github.com/windingtheropes/budget/based"
	"github.com/windingtheropes/budget/types"
)

var UserTable = based.NewTable[types.User, types.UserForm]("usr")
var SessionTable = based.NewTable[types.Session, types.SessionForm]("usr")

func GenToken(length int) string {
	var keyBytes = make([]byte, length)
	rand.Read(keyBytes)
	return b64.RawStdEncoding.EncodeToString(keyBytes)
}
