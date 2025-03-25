package auth

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