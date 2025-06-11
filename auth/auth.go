package auth

import (
	"crypto/rand"
	b64 "encoding/base64"
)

func GenToken(length int) string {
	var keyBytes = make([]byte, length)
	rand.Read(keyBytes)
	return b64.RawStdEncoding.EncodeToString(keyBytes)
}
