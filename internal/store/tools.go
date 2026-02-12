package store

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func (user User) SelectUser(db *pgx.Conn, ctx context.Context) (User, error) {
	queryMsg := `
		SELECT login, password, fio, email, rights
		FROM users
		WHERE login = $1;
		`

	row := db.QueryRow(ctx, queryMsg, user.Login)

	err := row.Scan(&user.Password)
	if err != nil {
		return user, err
	}

	return user, nil
}

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