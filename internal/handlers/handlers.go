package handlers

import (
	"Sybersports/internal/models"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"time"
)

type Service interface {
	CreateUser(ctx context.Context, user models.User) (models.User, error)
	GetUser(ctx context.Context, id int) (models.User, error)
}

type Handler struct {
	svc Service
}

func NewHandler(svc Service) *Handler {
	return &Handler{svc: svc}
}

func (h *Handler) RegistrationUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	ctx, cancel := context.WithTimeout(ctx, 20*time.Second)
	defer cancel()

	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method (POST only)", http.StatusMethodNotAllowed)
		return
	}

	var user models.User

	u, err := h.svc.CreateUser(ctx, user)
	if err != nil {
		switch {
		case errors.Is(err, errors.New("Invalid input")):
			http.Error(w, err.Error(), http.StatusBadRequest)
		default:
			http.Error(w, "internal error", http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(u)
}
