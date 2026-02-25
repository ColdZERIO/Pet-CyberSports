package repository

import (
	"Sybersports/internal/models"
	service "Sybersports/internal/service/storage"
	"context"
	"database/sql"
	"errors"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) SelectPostgres(ctx context.Context, id int) (models.User, error) {
	var user models.User

	queryMsg := `
		SELECT login, password, fio, email, rights
		FROM users
		WHERE login = $1;
		`

	row := r.db.QueryRowContext(ctx, queryMsg, user.Login)
	err := row.Scan(&user.Login, &user.Password, &user.FIO, &user.Email, &user.Rights)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.User{}, service.ErrUserNotFound
		}
		return models.User{}, err
	}
	return user, nil
}

func (r *Repository) CheckUserPostgres(ctx context.Context, user models.User) error {
	queryMsg := `
	SELECT login, email
	FROM users
	WHERE login = $1`

	row := r.db.QueryRowContext(ctx, queryMsg, user.Login)
	err := row.Scan(&user.Login, &user.Email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return service.ErrUserNotFound
		}
		return err
	}
	return nil
}

func (r *Repository) InsertPostgres(ctx context.Context, user models.User) (models.User, error) {
	queryMsg := `
		INSERT INTO users
		(login, password, fio, email, rights)
		VALUES ($1, $2, $3, $4, $5);
		`

	_, err := r.db.ExecContext(
		ctx,
		queryMsg,
		user.Login,
		user.Password,
		user.FIO,
		user.Email,
		user.Rights,
	)

	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (r *Repository) UpdatePostgres(ctx context.Context, user models.User) (models.User, error) {
	queryMsg := `
	UPDATE users
	SET login = $1, fio = $2, email = $3
	WHERE id = $4
	RETURNING id, login, fio, email;
	`

	var updated models.User
	err := r.db.QueryRowContext(
		ctx,
		queryMsg,
		user.Login,
		user.FIO,
		user.Email,
	).Scan(
		&updated.ID,
		&updated.Login,
		&updated.FIO,
		&updated.Email,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.User{}, service.ErrUserNotFound
		}
		return models.User{}, err
	}

	return updated, nil
}

func (r *Repository) DeletePostgres(ctx context.Context, id int) error {
	queryMsg := `
	DELETE FROM users
	WHERE id = $1;
	`

	res, err := r.db.ExecContext(ctx, queryMsg, id)
	if err != nil {
		return err
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return service.ErrUserNotFound
	}

	return nil
}
