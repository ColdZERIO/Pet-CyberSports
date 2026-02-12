package store

type User struct {
	ID       int    `json:"id"`
	Login    string `json:"login" validate:"required,min=6,max=20"`
	Password string `json:"password" validate:"required,min=6"`
	FIO      string `json:"fio"`
	Email    string `json:"email" validate:"required,email"`
	Rights   int    `json:"isadmin"`
}
