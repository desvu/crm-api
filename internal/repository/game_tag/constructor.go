package game_tag

import (
	"github.com/qilin/crm-api/internal/domain/repository"
	"github.com/qilin/crm-api/internal/env"
	"github.com/qilin/crm-api/internal/repository/game_tag/postgres"
)

func New(env *env.Store) repository.GameTagRepository {
	return postgres.New(env.Postgres)
}
