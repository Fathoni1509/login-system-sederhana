package dto

type RegisterUser struct {
	Name     string
	Email    string
	Phone    string
	Password string
}

type LoginUser struct {
	Email    string
	Password string
}
