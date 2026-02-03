package postgres

import (
	"context"
	"log"
)

func main() {
	ctx := context.Background()

	conn, err := CreateConnection(ctx)
	if err != nil {
		panic(err)
	}

	err = CreateTables(conn, ctx)
	if err != nil {
		log.Println("Bad database request")
		return
	}
}
