package repository

import (
	"github.com/badprinter/remember/internal/models"
)

func (r *Repository) InsertTask(t *models.Task, u *models.User) error {
	exeStr := "INSERT INTO tasks(user_id, document) VALUES ($1, $2) RETURNING id"
	err := r.SQL.QueryRow(exeStr, u.ID, t.Document).Scan(&t.ID)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) UpdateTask(t *models.Task) error {
	exeStr := "UPDATE tasks SET Document WHERE id = $1"
	_, err := r.SQL.Exec(exeStr, t.Document)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) DeleteTask(t *models.Task) error {

	_, err := r.SQL.Exec("DELETE FROM tasks WHERE id = $1", t.ID)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) GetTask(t *models.Task) error {
	queryStr := "SELECT user_id, document FROM tasks WHERE id = $1"
	err := r.SQL.QueryRow(queryStr, t.ID).Scan(&t.USER_ID, &t.Document)
	if err != nil {
		return err
	}
	return nil
}
