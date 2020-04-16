package game_revision_feature

import (
	"github.com/qilin/crm-api/internal/domain/repository"
	"github.com/qilin/crm-api/internal/env"
	"github.com/qilin/crm-api/internal/repository/game_revision_feature/postgres"
)

func New(env *env.Env) repository.GameRevisionFeatureRepository {
	return postgres.New(env.Store.Postgres)
}
