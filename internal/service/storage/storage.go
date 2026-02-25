package service

import (
	"Sybersports/internal/models"
	"Sybersports/internal/service/secure"
	"context"
	"errors"
	"strings"
)

// Надо добавить еще функций согласно интерфейсу

var ErrUserNotFound = errors.New("user not found")
var ErrInvalidInput = errors.New("invalid input")

type Repository interface {
	SelectPostgres(ctx context.Context, id int) (models.User, error)
	InsertPostgres(ctx context.Context, user models.User) (models.User, error)
	UpdatePostgres(ctx context.Context, user models.User) (models.User, error)
	DeletePostgres(ctx context.Context, id int) error
	CheckUserPostgres(ctx context.Context, user models.User) error
}

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) CreateUser(ctx context.Context, user models.User) (models.User, error) {
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	user.Login = strings.TrimSpace(user.Login)

	if user.Email == "" || user.Login == "" {
		return models.User{}, ErrInvalidInput
	}

	err := s.repo.CheckUserPostgres(ctx, user)
	if err != nil {
		return models.User{}, errors.New("user with login or email already exists")
	}

	newUser := models.User{
		Login:    user.Login,
		Password: user.Password,
		FIO:      user.FIO,
		Email:    user.Email,
	}

	pass, err := secure.HashPassword(user.Password, secure.DefaultParams)
	if err != nil {
		return models.User{}, errors.New("cant hash password")
	}

	newUser.Password = pass

	return s.repo.InsertPostgres(ctx, newUser)
}

func (s *Service) GetUser(ctx context.Context, id int) (models.User, error) {
	if id <= 0 {
		return models.User{}, ErrInvalidInput
	}

	return s.repo.SelectPostgres(ctx, id)
}
