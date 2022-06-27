package config

type RepositoryCfg struct {
	HOST     string `toml:"ip"`
	PORT     string `toml:"port"`
	USER     string `toml:"user_name"`
	PASSWORD string `toml:"password"`
	DB_NAME  string `toml:"data_base_name"`
	SSL_MODE string `toml:"ssl_mode"`
}
