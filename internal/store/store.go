package store

import (
	"Sybersports/internal/models"
	"context"
)

type UserStore interface {
	SelectUser(ctx context.Context, login string) (*models.User, error)
	InsertUser(ctx context.Context, user *models.User) error
	UpdateUser(ctx context.Context, user *models.User) error
	DeleteUser(ctx context.Context, id int) error
}
