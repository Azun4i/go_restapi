package apiserver

type Config struct {
	BindAddr    string `toml:"bind_addr"`
	DatabaseUrl string `toml:"database_url"`
	LogLevel    string `toml:"log_level"`
}

func NewConfig() *Config {
	return &Config{
		BindAddr: ":8080",
		LogLevel: "debug",
	}
}
