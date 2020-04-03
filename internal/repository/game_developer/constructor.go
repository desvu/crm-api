package game_developer

import (
	"github.com/qilin/crm-api/internal/domain/repository"
	"github.com/qilin/crm-api/internal/env"
	"github.com/qilin/crm-api/internal/repository/game_developer/postgres"
)

func New(env *env.Store) repository.GameDeveloperRepository {
	return postgres.New(env.Postgres)
}
