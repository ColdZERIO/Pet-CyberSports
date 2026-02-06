package pkg

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5"
)

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
