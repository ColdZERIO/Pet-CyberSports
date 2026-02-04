package postgres

import (
	"context"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	ctx := context.Background()

	if err := godotenv.Load(); err != nil {
		log.Println("Error read .env file")
		return
	}

	Conn, err := CreateConnection(ctx)
	if err != nil {
		log.Println("Error connection database")
		return
	}

	err = CreateTables(Conn, ctx)
	if err != nil {
		log.Println("Bad database request")
		return
	}
}
