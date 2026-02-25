package main

import (
	"Sybersports/internal/handlers"
	"Sybersports/internal/repository"
	service "Sybersports/internal/service/storage"
	postgres "Sybersports/pgk/postgresql"
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println(".env is not found")
		return
	}

	ctx := context.Background()
	defer ctx.Done()

	host := getEnv("DB_HOST", "")
	port := getEnv("DB_PORT", "")
	user := getEnv("DB_USER", "")
	password := getEnv("DB_PASSWORD", "")
	dbname := getEnv("DB_NAME", "")

	msgConn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", user, password, host, port, dbname)
	db, err := sql.Open("pgx", msgConn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Временно
	conn, err := postgres.CreateConnection(ctx)
	if err != nil {
		log.Fatal(err)
	}

	err = postgres.CreateTables(conn, ctx)
	if err != nil {
		log.Fatal(err)
	}

	repo := repository.NewRepository(db)
	svc := service.NewService(repo)
	hand := handlers.NewHandler(svc)

	srv := chi.NewRouter()
	srv.Post("/reg", hand.RegistrationUser)
	log.Println("Server START")

	err = http.ListenAndServe(":8080", srv)
	if err != nil {
		log.Fatal(err)
		return
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
