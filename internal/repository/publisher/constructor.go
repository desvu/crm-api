package publisher

import (
	"github.com/qilin/crm-api/internal/domain/repository"
	"github.com/qilin/crm-api/internal/env"
	"github.com/qilin/crm-api/internal/repository/publisher/postgres"
)

func New(env *env.Env) repository.PublisherRepository {
	return postgres.New(env.Store.Postgres)
}
