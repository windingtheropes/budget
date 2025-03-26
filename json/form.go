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
type TagForm struct {
	Token string `json:"token" binding:"required"`
	Name string `json:"name" bindings:"required"`
}
type ValueForm struct {
	Value string `json:"value" binding:"required"`
}