package handlers

import (
	"Sybersports/internal/models"
	"Sybersports/internal/service"
	"net/http"
)

func Auth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method (POST only)", http.StatusMethodNotAllowed)
		return
	}

	user := models.User{}

	user.Login = r.FormValue("login")
	user.Password = r.FormValue("password")

	service.VerifyPassword(user.Password, "")

	w.Write([]byte("Well CAM!"))
}
