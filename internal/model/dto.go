package model

type RegisterDTO struct {
	Name     string `form:"name"`
	Email    string `form:"email"`
	Password string `form:"password"`
}

type LoginDTO struct {
	Email    string `form:"email"`
	Password string `form:"password"`
}
