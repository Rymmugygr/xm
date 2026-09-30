package postgres

import (
	"fmt"

	"github.com/Rymmugygr/xm/config"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Postgres struct {
	Gorm *gorm.DB
}

func New(_ *logrus.Logger, cfg config.Db) (*Postgres, error) {
	db, err := gorm.Open(postgres.Open(cfg.Dsn), &gorm.Config{
		TranslateError: true,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	return &Postgres{Gorm: db}, nil
}

func (p *Postgres) Close() {
	return
}
