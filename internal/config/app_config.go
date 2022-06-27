package config

import (
	"fmt"
	"log"

	"github.com/BurntSushi/toml"
)

type ConfigApp struct {
	Repository *RepositoryCfg `toml:"db"`
	HttpServer *HttpServerCfg `toml:"http"`
}

func NewConfig(pathToCofingTOML string) *ConfigApp {
	var cfg ConfigApp
	_, err := toml.DecodeFile(pathToCofingTOML, &cfg)
	if err != nil {
		log.Fatalln("Fatal NweConfgi", err)
	}
	return &cfg
}

func (c *ConfigApp) String() string {
	return fmt.Sprintf(
		"\n%s\n%s\n%s\n%s\n%s\n%s\n%s\n%s\n%s\n%s\n%s\n%s\n",
		"--------------------",
		"SQL server",
		c.Repository.DB_NAME,
		c.Repository.HOST,
		c.Repository.PASSWORD,
		c.Repository.USER,
		c.Repository.PORT,
		c.Repository.SSL_MODE,
		"--------------------",
		"HTTP",
		c.HttpServer.Ip,
		c.HttpServer.Port,
	)
}
