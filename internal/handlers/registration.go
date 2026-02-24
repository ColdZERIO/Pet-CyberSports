package handlers

import (
	"Sybersports/internal/models"
	service "Sybersports/internal/service/storage"
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

	user := &models.User{}

	user, err := service.AddNewUser(r, ctx, user)
	if err != nil {
		http.Error(w, "Error add user to Database", http.StatusBadRequest)
		return
	}

	w.Write([]byte("New user " + user.FIO + "\nregistration success!"))
}
