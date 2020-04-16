package game_revision_tag

import (
	"github.com/qilin/crm-api/internal/domain/repository"
	"github.com/qilin/crm-api/internal/env"
	"github.com/qilin/crm-api/internal/repository/game_revision_tag/postgres"
)

func New(env *env.Env) repository.GameRevisionTagRepository {
	return postgres.New(env.Store.Postgres)
}
