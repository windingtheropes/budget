package argent

type User struct {
	Id int
	Name string
	Email string
	Password string
}

type Session struct {
	Id int
	Token string
	User_Id int
}

type Currency struct {
	Id string
}
type BudgetEntry struct {
	Id int
	User_Id int
	Amount float64
	Currency string
}
type TagAssignment struct {
	Id int
	Tag_Id int
	Entry_Id int
}
type Tag struct {
	Id int
	Name int
	User_Id int
}