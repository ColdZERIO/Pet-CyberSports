package store

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func (user User) DeleteUser(db *pgx.Conn,ctx context.Context) error {
	queryMsg := `
	DELETE FROM users
	WHERE id = $1;
	`

	_, err := db.Exec(ctx, queryMsg, user.ID)
	if err != nil {
		return err
	}

	return nil
}
