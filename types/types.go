package types

type UserIdentifier interface {
	Email | UserID
}

type UserID int64
type Email string
type User struct {
	Id       int64
	First_Name     string
	Last_Name string
	Email    string
	Password string
}
type UserForm struct {
	First_Name     string `db:"first_name"`
	Last_Name string `db:"last_name"`
	Email    string `db:"email"`
	Password string `db:"pass"`
}
type Session struct {
	Id      int64
	Token   string
	User_Id int64
	Expiry  int64
}
type SessionForm struct {
	Token   string `db:"token"`
	User_Id int64 `db:"user_id"`
	Expiry  int64 `db:"expiry"`
}
type Currency struct {
	Id string
}
type TransactionType struct {
	Id   int64
	Name string
	Positive bool
}
type TransactionTypeForm struct {
	Name string `db:"type_name"`
	Positive bool `db:"positive"`
}
type TransactionEntry struct {
	Id             int64
	User_Id        int64
	Type_Id        int64
	Msg            string
	Amount         float64
	Currency       string
	Unix_Timestamp int64
	Vendor         string
}
type TransactionEntryForm struct {
	User_Id        int64	`db:"user_id"`
	Type_Id        int64	`db:"type_id"`
	Msg            string	`db:"msg"`
	Amount         float64	`db:"amount"`
	Currency       string	`db:"currency"`
	Unix_Timestamp int64	`db:"unix_timestamp"`
	Vendor         string	`db:"vendor"`
}
type HydTransactionEntry struct {
	Id             int64
	User_Id        int64
	Type_Id        int64
	Msg            string
	Amount         float64
	Currency       string
	Tags           []Tag
	Unix_Timestamp int64
	Vendor         string
	Budget_Entries []BudgetEntry
}
type TagAssignment struct {
	Id             int64
	Tag_Id         int64
	Transaction_Id int64
}
type TagAssignmentForm struct {
	Tag_Id         int64 `db:"tag_id"`
	Transaction_Id int64 `db:"entry_id"`
}

type TagID int64
type TagName string
type Tag struct {
	Id   int64
	Name string
}
type TagForm struct {
	Name        string   `db:"tag_name"`
}
type HydTag struct {
	Id         int64
	User_Id    int64
	Name       string
	Tag_Budgets []TagBudget
}
type HydBudget struct {
	Id         int64
	Name       string
	Goal 	   float64
	Tag_Budgets []TagBudget
}
type TagOwnership struct {
	Id      int64
	Tag_Id  int64
	User_Id int64
}
type TagOwnershipForm struct {
	Tag_Id  int64 `db:"tag_id"`
	User_Id int64 `db:"user_id"`
}

type Budget struct {
	Id      int64
	User_Id int64
	Name    string
	Goal    float64
}
type BudgetForm struct {
	User_Id int64 `db:"user_id"`
	Name    string `db:"budget_name"`
	Goal    float64 `db:"goal"`
}

type BudgetEntry struct {
	Id             int64
	Transaction_Id int64
	Budget_Id      int64
	Amount         float64
}
type BudgetEntryForm struct {
	Transaction_Id int64 `db:"transaction_id"`
	Budget_Id      int64 `db:"budget_id"`
	Amount         float64 `db:"amount"`
}

type TagBudget struct {
	Id        int64
	Tag_Id    int64
	Budget_Id int64
	Goal      float64
	Type_Id   int64
}
type TagBudgetForm struct {
	Tag_Id    int64 `db:"tag_id"`
	Budget_Id int64 `db:"budget_id"`
	Goal      float64 `db:"goal"`
	Type_Id   int64 `db:"type_id`
}
