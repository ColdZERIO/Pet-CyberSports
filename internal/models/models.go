package models

type User struct {
	ID       int    `json:"id"`
	Login    string `json:"login"`
	Password string `json:"password"`
	FIO      string `json:"fio"`
	Email    string `json:"email"`
}
