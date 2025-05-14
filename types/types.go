package types

type UserIdentifier interface {
	Email | UserID
}
type TagIdentifier interface {
	TagID | UserID | TagName
}

type UserID int
type Email string
type User struct {
	Id       int
	Name     string
	Email    string
	Password string
}
type Session struct {
	Id      int
	Token   string
	User_Id int
	Expiry  int64
}
type Currency struct {
	Id string
}
type TransactionType struct {
	Id   int
	Name string
}
type TransactionEntry struct {
	Id             int
	User_Id        int
	Type_Id        int
	Msg            string
	Amount         float64
	Currency       string
	Unix_Timestamp int64
	Vendor         string
}
type HydTransactionEntry struct {
	Id             int
	User_Id        int
	Type_Id        int
	Msg            string
	Amount         float64
	Currency       string
	Tags           []Tag
	Unix_Timestamp int64
	Vendor         string
}
type TagAssignment struct {
	Id             int
	Tag_Id         int
	Transaction_Id int
}

type TagID int
type TagName string
type Tag struct {
	Id   int
	Name string
}
type HydTag struct {
	Id 			int
	User_Id 	int
	Name 		string
	Budget_Id	[]int
}
type TagOwnership struct {
	Id      int
	Tag_Id  int
	User_Id int
}

type Budget struct {
	Id      int
	User_Id int
	Name    string
	Type_Id int
	Goal    float64
}

type BudgetEntry struct {
	Id             int
	Transaction_Id int
	Budget_Id      int
	Amount         float64
}

type TagBudget struct {
	Id        int
	Tag_Id    int
	Budget_Id int
	Goal      float64
	Type_Id   int
}
