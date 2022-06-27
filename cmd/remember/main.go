package main

import (
	"log"

	"github.com/badprinter/remember/internal/api"
	"github.com/badprinter/remember/internal/config"
)

var (
	FileConfig = "./config/config.toml"
)

func main() {
	cfg := config.NewConfig(FileConfig)
	log.Println(cfg)
	app := api.NewApp(cfg)

	app.Http.Start()
}
