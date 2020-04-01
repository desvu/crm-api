package game

import (
	"github.com/qilin/crm-api/internal/domain/repository"
	"github.com/qilin/crm-api/internal/env"
	"github.com/qilin/crm-api/internal/repository/game/postgres"
)

func New(env *env.Store) repository.IGameRepository {
	return postgres.New(env.Postgres)
}
