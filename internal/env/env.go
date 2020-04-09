package env

import (
	"context"
	"errors"

	"github.com/qilin/crm-api/internal/config"
	"github.com/qilin/crm-api/internal/env/migration/postgres"
	"github.com/qilin/crm-api/pkg/transactor"
)

type Env struct {
	Store *Store
}

func New(ctx context.Context, cfg *config.Config, transactionStore *transactor.Store) (*Env, error) {
	if cfg == nil {
		return nil, errors.New("config is nil")
	}

	storeEnv, err := newStore(ctx, cfg.Store, transactionStore)
	if err != nil {
		return nil, err
	}

	if err = postgres.Migrate(storeEnv.Postgres.Handler.GetConnection()); err != nil {
		return nil, err
	}

	return &Env{
		Store: storeEnv,
	}, nil
}
