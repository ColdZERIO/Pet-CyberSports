package store

import (
	"Sybersports/internal/models"
	"context"

	"github.com/jackc/pgx/v5"
)

func InsertDB(db *pgx.Conn, user models.User, ctx context.Context) error {
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
