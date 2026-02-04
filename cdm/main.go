package main

import (
	"Sybersports/internal/handlers"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	srv := chi.NewRouter()
	log.Println("Srever START")

	srv.Post("/registration", handlers.Registration)
	srv.Post("/auth", handlers.Auth)

	err := http.ListenAndServe(":8080", srv)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("Server STOP")
}
