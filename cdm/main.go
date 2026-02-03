package main

import (
	"log"
	"net/http"
	"school/internal/handlers"

	"github.com/go-chi/chi/v5"
)

func main() {
	srv := chi.NewRouter()
	log.Println("Srever START")

	srv.Post("/registration", handlers.Registration)
	srv.Post("/auth", handlers.Auth)

	log.Fatal(http.ListenAndServeTLS(":8080", "server.crt", "server.key", srv))

	err := http.ListenAndServe(":8080", srv)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("Server STOP")
}
