package types
import "time"
func (session *Session) IsValid() bool {
	if time.Now().Unix() <= session.Expiry {
		return true
	}
	return false
}