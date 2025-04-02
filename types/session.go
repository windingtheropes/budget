package types
import "time"
func (session *Session) IsValid() bool {
	return time.Now().Unix() <= session.Expiry
}