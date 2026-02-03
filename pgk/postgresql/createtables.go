package postgres

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func CreateTables(conn *pgx.Conn, ctx context.Context) error {
	msg := `
	CREATE TABLE users IF NOT EXISTS (
		id SERIAL PRIMARY KEY NOT NULL,
		login VARCHAR(20) NOT NUL,
		password VARCHAR(200) NOT NUL,
		fio VARCHAR(60) NOT NULL,
		email VARCHAR(60),

		UNIQUE(login)
	)`

	_, err := conn.Exec(ctx, msg)
	return err
}
