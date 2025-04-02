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
type TransactionEntry struct {
	Id             int
	User_Id        int
	Amount         float64
	Currency       string
	Unix_Timestamp int64
}
type HydTransactionEntry struct {
	Id int
	User_Id int
	Amount float64
	Currency string
	Tags []Tag
	Unix_Timestamp int64
}
type TagAssignment struct {
	Id       int
	Tag_Id   int
	Entry_Id int
}

type TagID int
type TagName string
type Tag struct {
	Id      int
	Name    string
	User_Id int
}
