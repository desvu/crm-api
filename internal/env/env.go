package env

import (
	"github.com/isayme/go-amqp-reconnect/rabbitmq"
	"github.com/qilin/crm-api/internal/config"
	"github.com/qilin/crm-api/internal/env/migration/postgres"
	"github.com/qilin/crm-api/pkg/transactor"
)

type Env struct {
	Store  *Store
	Rabbit *rabbitmq.Connection
}

func New(transactor *transactor.Transactor) (*Env, error) {
	cfg, err := config.New()
	if cfg == nil {
		return nil, err
	}

	rabbitEnv, err := newRabbit()
	if err != nil {
		return nil, err
	}

	storeEnv, err := newStore(cfg.Store, transactor.GetStore())
	if err != nil {
		return nil, err
	}

	if err = postgres.Migrate(storeEnv.Postgres.Handler.GetConnection()); err != nil {
		return nil, err
	}

	return &Env{
		Store:  storeEnv,
		Rabbit: rabbitEnv,
	}, nil
}
