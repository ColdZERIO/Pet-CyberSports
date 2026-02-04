package postgres

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
)

type ParamsDB struct {
	host     string
	port     string
	user     string
	password string
	dbname   string
}

func CreateConnection(ctx context.Context) (*pgx.Conn, error) {
	params := ParamsDB{
		host:     getEnv("DB_HOST", "localhost"),
		port:     getEnv("DB_PORT", "5432"),
		user:     getEnv("DB_USER", "postgres"),
		password: getEnv("DB_PASSWORD", "123"),
		dbname:   getEnv("DB_NAME", "cybersports"),
	}

	connMsg := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", params.user, params.password, params.host, params.port, params.dbname)

	conn, err := pgx.Connect(ctx, connMsg)
	if err != nil {
		log.Println("Error connection database")
		return nil, err
	}

	err = conn.Ping(ctx)
	if err != nil {
		log.Println("Error ping database")
		return nil, err
	}

	return conn, nil
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
