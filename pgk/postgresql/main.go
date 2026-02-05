package postgres

import (
	"context"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println(".env is not found")
		return
	}

	ctx := context.Background()

	conn, err := CreateConnection(ctx)
	if err != nil {
		log.Fatal("Database connection failed")
		return
	}

	err = CreateTables(conn, ctx)
	if err != nil {
		log.Fatal("Database create failed")
		return
	}
}
