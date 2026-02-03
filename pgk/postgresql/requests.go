package postgres

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func SelectID(db *pgx.Conn, ctx context.Context) 