package developer

import (
	"github.com/qilin/crm-api/internal/domain/repository"
	"github.com/qilin/crm-api/internal/env"
	"github.com/qilin/crm-api/internal/repository/developer/postgres"
)

func New(env *env.Store) repository.DeveloperRepository {
	return postgres.New(env.Postgres)
}
