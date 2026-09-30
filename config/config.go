package config

import (
	"os"
)

type (
	Config struct {
		App  App
		Http Http
		Db   Db
	}

	App struct {
		Name    string
		Version string
	}
	Http struct {
		Addr string
	}
	Db struct {
		Dsn string
	}
)

// NewConfig returns app config.
func NewConfig() (*Config, error) {
	cfg := &Config{
		App: App{
			Name:    os.Getenv("APP_NAME"),
			Version: os.Getenv("APP_VERSION"),
		},
		Http: Http{
			Addr: os.Getenv("LISTEN_ADDR_HTTP"),
		},
		Db: Db{
			Dsn: os.Getenv("PG_DSN"),
		},
	}

	return cfg, nil
}
