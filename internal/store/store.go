package store

import "context"

type UserStore interface {
	SelectUser(ctx context.Context, login string) (*User, error)
	InsertUser(ctx context.Context, user *User) error
	UpdateUser(ctx context.Context, user *User) error
	DeleteUser(ctx context.Context, id int) error
}
