package game_feature

import (
	"github.com/qilin/crm-api/internal/domain/repository"
	"github.com/qilin/crm-api/internal/env"
	"github.com/qilin/crm-api/internal/repository/game_feature/postgres"
)

func New(env *env.Store) repository.GameFeatureRepository {
	return postgres.New(env.Postgres)
}
