package app

import (
	"errors"

	"github.com/qilin/crm-api/internal/config"
)

type App struct {
	StorageURL string
}

func New(cfg *config.Config) (*App, error) {
	storageUrl := cfg.App.StorageURL
	if storageUrl == "" {
		return nil, errors.New("storage url must be set")
	}

	if len(storageUrl) > 0 && storageUrl[:len(storageUrl)-1] == "/" {
		storageUrl = storageUrl[:len(storageUrl)-1]
	}

	return &App{
		StorageURL: storageUrl,
	}, nil
}
