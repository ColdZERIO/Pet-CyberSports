package service

import (
	"Sybersports/internal/models"
	"database/sql"
)

func Password(db *sql.DB, user models.User) {
	row := db.QueryRow("SELECT password FROM users WHERE id = $1", user.ID)

	row.Scan(&user.Password)
}
