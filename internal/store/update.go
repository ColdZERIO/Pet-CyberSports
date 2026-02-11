package store

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func (user User) UpdateUser(db *pgx.Conn, ctx context.Context) error {
	queryMsg := `
	UPDATE users
	SET rights = $1
	WHERE id = $2;
	`

	_, err := db.Exec(
		ctx,
		queryMsg,
		user.Login,
		user.Password,
		user.FIO,
		user.Email,
	)

	if err != nil {
		return err
	}

	return nil
}
