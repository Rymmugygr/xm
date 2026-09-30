package config

import (
	"os"
)

type (
	Config struct {
		App  App
		Http Http
		Db   Db
		Jwt  Jwt
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
	Jwt struct {
		Provider         string
		Secret           string
		Algorithm        string
		GetTokenUri      string
		ValidateTokenUri string
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
		Jwt: Jwt{
			Secret: os.Getenv("JWT_SECRET"),
		},
	}

	return cfg, nil
}
