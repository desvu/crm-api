package genre

import (
	"github.com/qilin/crm-api/internal/domain/repository"
	"github.com/qilin/crm-api/internal/env"
	"github.com/qilin/crm-api/internal/repository/genre/postgres"
)

func New(env *env.Env) repository.GenreRepository {
	return postgres.New(env.Store.Postgres)
}
