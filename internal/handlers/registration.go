package handlers

import (
	"Sybersports/internal/models"
	"Sybersports/internal/service"
	"Sybersports/internal/store"
	postgres "Sybersports/pgk/postgresql"
	"context"
	"net/http"
	"time"
)

func Registration(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()
	ctx, cancel := context.WithTimeout(ctx, 20*time.Second)
	defer cancel()

	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method (POST only)", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Invalid form", http.StatusBadRequest)
		return
	}

	user := models.User{
		Login:    r.FormValue("login"),
		Password: r.FormValue("password"),
		FIO:      r.FormValue("fio"),
		Email:    r.FormValue("email"),
	}

	user.Password, err = service.HashPassword(user.Password, service.DefaultParams)
	if err != nil {
		http.Error(w, "Error with password", http.StatusBadRequest)
	}

	conn, err := postgres.CreateConnection(ctx)
	if err != nil {
		http.Error(w, "Error connection Database", http.StatusBadRequest)
		return
	}
	err = store.InsertDB(conn, user, ctx)

	if err != nil {
		http.Error(w, "Error add data to Database", http.StatusBadRequest)
		return
	}

	w.Write([]byte("New user " + user.FIO + "\nregistration success!"))
}
