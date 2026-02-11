package store

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func (user User) SelectUser(db *pgx.Conn, ctx context.Context) (string, error) {
	queryMsg := `
		SELECT password
		FROM users
		WHERE login = $1;
		`

	row := db.QueryRow(ctx, queryMsg, user.ID)

	err := row.Scan(&user.Password)
	if err != nil {
		return "", err
	}

	return user.Password, nil
}
