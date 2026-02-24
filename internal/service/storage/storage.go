package service

import (
	"Sybersports/internal/models"
	service "Sybersports/internal/service/secure"
	postgres "Sybersports/pgk/postgresql"
	"context"
	"errors"
	"net/http"
)

func AddNewUser(r *http.Request, ctx context.Context, user *models.User) (*models.User, error) {
	db, err := postgres.CreateConnection(ctx)
	if err != nil {
		return nil, errors.New()
	}

	err = r.ParseForm()
	if err != nil {
		return nil, errors.New("cant parse user form")
	}

	login := r.FormValue("login")
	password := r.FormValue("password")
	fio := r.FormValue("fio")
	email := r.FormValue("email")

	password, err = service.HashPassword(user.Password, service.DefaultParams)
	if err != nil {
		return nil, errors.New("problems with hash password")
	}

	user.Login = login
	user.Password = password
	user.FIO = fio
	user.Email = email



	return user, nil
}
