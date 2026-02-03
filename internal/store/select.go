package store

import (
	"Sybersports/internal/models"
	"context"

	"github.com/jackc/pgx/v5"
)

func SelectPasswordDB(db *pgx.Conn, user models.User, ctx context.Context) (string, error) {
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

func UpdateUserDB(db *pgx.Conn, user models.User, ctx context.Context) error {
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

func DeleteDB(db *pgx.Conn, user models.User, ctx context.Context) error {
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
