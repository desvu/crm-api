package app

import (
	"errors"

	"github.com/qilin/crm-api/internal/config"
)

type App struct {
	StorageURL string
}

func New(cfg *config.Config) (*App, error) {
	if cfg.App.StorageURL == "" {
		return nil, errors.New("storage url must be set")
	}

	return &App{
		StorageURL: cfg.App.StorageURL,
	}, nil
}
