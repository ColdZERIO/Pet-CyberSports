package store

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func (user User) InsertUser(db *pgx.Conn, ctx context.Context) error {
	queryMsg := `
		INSERT INTO users
		(login, password, fio, email, rights)
		VALUES ($1, $2, $3, $4, $5);
		`

	_, err := db.Exec(
		ctx,
		queryMsg,
		user.Login,
		user.Password,
		user.FIO,
		user.Email, // add params
	)

	if err != nil {
		return err
	}

	return nil
}
