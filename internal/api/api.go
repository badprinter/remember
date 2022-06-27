package api

import (
	"net/http"

	"github.com/badprinter/remember/internal/config"
	"github.com/badprinter/remember/internal/httpserver"
	"github.com/badprinter/remember/internal/repository"
)

type App struct {
	Http *httpserver.ServerHTTP
	Repo *repository.Repository
}

func NewApp(cfg *config.ConfigApp) *App {

	app := &App{
		Http: httpserver.CreateHttpServer(cfg.HttpServer),
		Repo: repository.CreateRepository(cfg.Repository),
	}

	app.Http.Router.HandleFunc("/login", app.login).Methods(http.MethodPost)

	// Document
	app.Http.Router.HandleFunc("/document", app.Ducument).Methods(http.MethodPost)

	return app
}
