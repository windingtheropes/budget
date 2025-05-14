package types

type UserIdentifier interface {
	Email | UserID
}


type UserID int64
type Email string
type User struct {
	Id       int64
	Name     string
	Email    string
	Password string
}
type Session struct {
	Id      int64
	Token   string
	User_Id int64
	Expiry  int64
}
type Currency struct {
	Id string
}
type TransactionType struct {
	Id   int64
	Name string
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

type TagID int64
type TagName string
type Tag struct {
	Id   int64
	Name string
}
type HydTag struct {
	Id         int64
	User_Id    int64
	Name       string
	Tag_Budgets []TagBudget
}
type TagOwnership struct {
	Id      int64
	Tag_Id  int64
	User_Id int64
}

type Budget struct {
	Id      int64
	User_Id int64
	Name    string
	Type_Id int64
	Goal    float64
}

type BudgetEntry struct {
	Id             int64
	Transaction_Id int64
	Budget_Id      int64
	Amount         float64
}

type TagBudget struct {
	Id        int64
	Tag_Id    int64
	Budget_Id int64
	Goal      float64
	Type_Id   int64
}
