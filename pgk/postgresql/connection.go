package postgres

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func CreateConnection(ctx context.Context) (*pgx.Conn, error) {
	host := getEnv("DB_HOST", "")
	port := getEnv("DB_PORT", "")
	user := getEnv("DB_USER", "")
	password := getEnv("DB_PASSWORD", "")
	dbname := getEnv("DB_NAME", "")

	msgConn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", user, password, host, port, dbname)
	log.Println(msgConn)

	return pgx.Connect(ctx, msgConn)
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func CreateTables(conn *pgx.Conn, ctx context.Context) error {
	query := `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		login VARCHAR(20) NOT NULL,
		password VARCHAR(200) NOT NULL,
		fio VARCHAR(60) NOT NULL,
		email VARCHAR(60),
		rights INTEGER NOT NULL DEFAULT 1,
		UNIQUE(login)
	);
	
	CREATE INDEX IF NOT EXISTS users_name ON users(fio);
	`

	_, err := conn.Exec(ctx, query)
	if err != nil {
		log.Println("Error create tables")
		return err
	}

	return nil
}
