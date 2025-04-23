package json

type NewAccountForm struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}
type LoginForm struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
type SessionForm struct {
	Token	string `json:"token"`
}
type NewTagForm struct {
	Name string `json:"name" bindings:"required"`
}
type ValueForm struct {
	Value string `json:"value" binding:"required"`
}
type NewTransactionForm struct {
	Msg string `json:"msg" binding:"required"`
	Amount   float64 `json:"amount" binding:"required"`
	Currency string  `json:"currency" binding:"required"`
	Tags []int `json:"tags" binding:"required"`
	Unix_Timestamp int `json:"unix_timestamp" binding:"required"`
}
