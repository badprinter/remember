package api

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/badprinter/remember/internal/models"
)

// TODO Document нужно ассоциировать с юзером
func (a *App) Ducument(w http.ResponseWriter, r *http.Request) {
	text, err := ioutil.ReadAll(r.Body)
	if err != nil {
		// TODO
		return
	}

	// <-------------
	var t *models.Task
	t.Document = string(text)
	u, err := a.Repo.GetUser(123132) // <---
	if err != nil {
		return
	}
	err = a.Repo.InsertTask(t, u)
	if err != nil {
		log.Println("Документ,", err)
	}
	// <------------
}
