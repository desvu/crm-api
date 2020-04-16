package game_revision_developer

import (
	"github.com/qilin/crm-api/internal/domain/repository"
	"github.com/qilin/crm-api/internal/env"
	"github.com/qilin/crm-api/internal/repository/game_revision_developer/postgres"
)

func New(env *env.Env) repository.GameRevisionDeveloperRepository {
	return postgres.New(env.Store.Postgres)
}
