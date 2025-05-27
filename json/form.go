package json

type AccountForm struct {
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
type TagForm struct {
	Name        string          `json:"name" bindings:"required"`
	Tag_Budgets []TagBudgetForm `json:"tag_budgets"`
}
type BudgetForm struct {
	Type_Id int64     `json:"type_id" binding:"required"`
	Name    string  `json:"name" binding:"required"`
	Goal    float64 `json:"goal" binding:"required"`
}
type TagBudgetForm struct {
	Budget_Id int64     `json:"budget_id" binding:"required"`
	Goal      float64 `json:"goal" binding:"required"`
	Type_Id   int64     `json:"type_id" binding:"required"`
}

type ValueForm[T any] struct {
	Value T `json:"value" binding:"required"`
}
type BudgetEntryForm struct {
	Budget_Id int64     `json:"budget_id" binding:"required"`
	Amount    float64 `json:"amount" binding:"required"`
}
type TransactionForm struct {
	Type_Id        int64     `json:"type_id" binding:"required"`
	Msg            string  `json:"msg"`
	Amount         float64 `json:"amount" binding:"required"`
	Currency       string  `json:"currency" binding:"required"`
	Tags           []int64   `json:"tags"`
	Unix_Timestamp int64     `json:"unix_timestamp" binding:"required"`
	Vendor         string  `json:"vendor"`
	Budget_Entries []BudgetEntryForm `json:"budget_entries"`
}
