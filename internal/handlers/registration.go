package handlers

import (
	"Sybersports/internal/models"
	"Sybersports/internal/service"
	"encoding/json"
	"log"
	"net/http"
	"os"
)

func Registration(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method (POST only)", http.StatusMethodNotAllowed)
		return
	}

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
		panic(err)
	}

	file, err := os.OpenFile("C:/GO/PracticeGO/base.json", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0755)
	if err != nil {
		log.Println("Error read the json file", err)
		return
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent(" ", "	")
	err = encoder.Encode(user)
	if err != nil {
		log.Println(err)
		return
	}

	w.Write([]byte("New user " + user.FIO + "\nregistration success!"))
}
