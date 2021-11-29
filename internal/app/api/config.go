package api

import "github.com/drTragger/powerfulAPI/storage"

// General instance for API server of REST application

type Config struct {
	BindAddr    string `toml:"bind_addr"` // Port
	LoggerLevel string `toml:"logger_level"`
	Storage     *storage.Config
}

func NewConfig() *Config {
	return &Config{
		BindAddr:    ":9001",
		LoggerLevel: "debug",
		Storage:     storage.NewConfig(),
	}
}
