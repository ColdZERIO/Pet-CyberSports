package store

import (
	"context"

	"github.com/jackc/pgx/v5"
)

type Postgres interface {
	SelectUser(*pgx.Conn, context.Context) (string, error)
	InsertUser(*pgx.Conn, context.Context) error
	UpdateUser(*pgx.Conn, context.Context) error
	DeleteUser(*pgx.Conn, context.Context) error
}

type User struct {
	ID       int    `json:"id"`
	Login    string `json:"login" validate:"required,min=6,max=20"`
	Password string `json:"password" validate:"required,min=6"`
	FIO      string `json:"fio"`
	Email    string `json:"email" validate:"required,email"`
	Rights   int    `json:"isadmin"`
}
