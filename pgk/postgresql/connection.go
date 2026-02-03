package postgres

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func CreateConnection(ctx context.Context) (*pgx.Conn, error) {
	conn, err := pgx.Connect(ctx, "postgres://postgres:123@localhost:5432/postgres")
	if err != nil {
		panic(err)
	}

	err = conn.Ping(ctx)
	if err != nil {
		panic(err)
	}

	return conn, nil
}
