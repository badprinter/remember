package repository

import (
	"github.com/badprinter/remember/internal/models"
)

func (r *Repository) InsertUser(u *models.User) error {
	err := r.SQL.QueryRow(
		"INSERT INTO users(username, password_encrypted) VALUES ($1, $2) RETURNING id",
		u.UserName, u.PasswordEncrypted,
	).Scan(&u.ID)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) UpdateUser(u *models.User) error {
	_, err := r.SQL.Exec(
		"UPDATE users SET username = $2, password_encrypted = $3) WHERE id=$1",
		u.ID, u.UserName, u.PasswordEncrypted,
	)
	if err != nil {
		return err
	}
	return nil
}

// GetUser first search User by UserName or Id than return all field, if not search return error
func (r *Repository) GetUser(id int) (*models.User, error) {
	var u models.User
	rowStr := "SELECT username, password_encrypted FROM users WHERE id = $1"
	err := r.SQL.QueryRow(
		rowStr,
		id,
	).Scan(&u.UserName, &u.PasswordEncrypted)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func (r *Repository) GetUserByUserName(userName string) *models.User {
	var u *models.User
	err := r.SQL.QueryRow("SELECT id, password_encrypted FROM users").Scan(u.ID, u.PasswordEncrypted)
	if err != nil {
		return nil
	}
	u.UserName = userName
	return u
}

func (r *Repository) DropUser(u *models.User) error {
	exe := "DELETE FROM users WHERE id = $1"
	_, err := r.SQL.Exec(exe, u.ID)
	if err != nil {
		return err
	}
	return nil
}
