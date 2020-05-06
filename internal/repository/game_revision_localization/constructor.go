package game_revision_localization

import (
	"github.com/qilin/crm-api/internal/domain/repository"
	"github.com/qilin/crm-api/internal/env"
	"github.com/qilin/crm-api/internal/repository/game_revision_localization/postgres"
)

func New(env *env.Env) repository.GameRevisionLocalizationRepository {
	return postgres.New(env.Store.Postgres)
}
