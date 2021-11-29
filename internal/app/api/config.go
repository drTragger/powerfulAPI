package api

// General instance for API server of REST application

type Config struct {
	BindAddr string `toml:"bind_addr"` // Port
}

func NewConfig() *Config {
	return &Config{
		BindAddr: ":9001",
	}
}
