package game_revision_genre

import (
	"github.com/qilin/crm-api/internal/domain/repository"
	"github.com/qilin/crm-api/internal/env"
	"github.com/qilin/crm-api/internal/repository/game_revision_genre/postgres"
)

func New(env *env.Env) repository.GameRevisionGenreRepository {
	return postgres.New(env.Store.Postgres)
}
