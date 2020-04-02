package tag

import (
	"github.com/qilin/crm-api/internal/domain/repository"
	"github.com/qilin/crm-api/internal/env"
	"github.com/qilin/crm-api/internal/repository/tag/postgres"
)

func New(env *env.Store) repository.TagRepository {
	return postgres.New(env.Postgres)
}
