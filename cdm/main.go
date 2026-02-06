package main

import (
	"Sybersports/internal/handlers"
	pkg "Sybersports/pgk/postgresql"
	"context"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println(".env is not found")
		return
	}

	ctx := context.Background()

	conn, err := pkg.CreateConnection(ctx)
	if err != nil {
		log.Fatal("Database connection failed")
		return
	}

	err = pkg.CreateTables(conn, ctx)
	if err != nil {
		log.Fatal("Database create failed")
		return
	}

	srv := chi.NewRouter()
	log.Println("Srever START")

	srv.Post("/registration", handlers.Registration)
	srv.Post("/auth", handlers.Auth)

	err = http.ListenAndServe(":8080", srv)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("Server STOP")
}
