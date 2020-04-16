package game_revision

import (
	"github.com/qilin/crm-api/internal/domain/repository"
	"github.com/qilin/crm-api/internal/env"
	"github.com/qilin/crm-api/internal/repository/game_revision/postgres"
)

func New(env *env.Env) repository.GameRevisionRepository {
	return postgres.New(env.Store.Postgres)
}
