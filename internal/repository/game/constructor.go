package game

import (
	"github.com/qilin/crm-api/internal/domain/repository"
	"github.com/qilin/crm-api/internal/env"
	"github.com/qilin/crm-api/internal/repository/game/postgres"
)

func New(env *env.Env) repository.GameRepository {
	return postgres.New(env.Store.Postgres)
}
