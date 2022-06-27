package httpserver

import (
	"github.com/badprinter/remember/internal/config"

	"github.com/gorilla/mux"
)

type ServerHTTP struct {
	Router *mux.Router
	Config *config.HttpServerCfg
}

func CreateHttpServer(cfg *config.HttpServerCfg) *ServerHTTP {
	return &ServerHTTP{
		Router: mux.NewRouter(),
		Config: cfg,
	}
}
