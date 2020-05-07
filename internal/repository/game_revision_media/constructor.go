package game_revision_media

import (
	"github.com/qilin/crm-api/internal/domain/repository"
	"github.com/qilin/crm-api/internal/env"
	"github.com/qilin/crm-api/internal/repository/game_revision_media/postgres"
)

func New(env *env.Env) repository.GameRevisionMediaRepository {
	return postgres.New(env.Store.Postgres)
}
