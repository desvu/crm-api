package feature

import (
	"github.com/qilin/crm-api/internal/domain/repository"
	"github.com/qilin/crm-api/internal/env"
	"github.com/qilin/crm-api/internal/repository/feature/postgres"
)

func New(env *env.Env) repository.FeatureRepository {
	return postgres.New(env.Store.Postgres)
}
