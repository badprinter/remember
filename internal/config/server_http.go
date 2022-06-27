package config

type HttpServerCfg struct {
	Ip   string `toml:"ip"`
	Port string `toml:"port"`
}
