package dto

type Auth struct {
	Email    string `form:"email"`
	Password string `form:"password"`
}