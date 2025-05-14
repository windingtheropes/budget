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
	Token string `json:"token"`
}
type NewTagForm struct {
	Name string `json:"name" bindings:"required"`
}
type NewBudgetForm struct {
	Type_Id        int     `json:"type_id" binding:"required"`
	Name            string  `json:"name" binding:"required"`
	Goal         float64 `json:"goal" binding:"required"`
}
type NewTagBudgetForm struct {
	Tag_Id    int  	`json:"tag_id" binding:"required"`
	Budget_Id int `json:"budget_id" binding:"required"`
	Goal      float64 `json:"goal" binding:"required"`
	Type_Id   int `json:"type_id" binding:"required"`
}

type ValueForm[T any] struct {
	Value T `json:"value" binding:"required"`
}
type NewBudgetEntryForm struct {
	Transaction_Id int	`json:"transaction_id" binding:"required"`
	Budget_Id 	   int	`json:"budget_id" binding:"required"`
	Amount		   float64  `json:"amount" binding:"required"`
}
type NewTransactionForm struct {
	Type_Id        int     `json:"type_id" binding:"required"`
	Msg            string  `json:"msg"`
	Amount         float64 `json:"amount" binding:"required"`
	Currency       string  `json:"currency" binding:"required"`
	Tags           []int   `json:"tags"`
	Unix_Timestamp int     `json:"unix_timestamp" binding:"required"`
	Vendor         string  `json:"vendor"`
}
