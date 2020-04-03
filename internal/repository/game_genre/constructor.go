package game_genre

import (
	"github.com/qilin/crm-api/internal/domain/repository"
	"github.com/qilin/crm-api/internal/env"
	"github.com/qilin/crm-api/internal/repository/game_genre/postgres"
)

func New(env *env.Store) repository.GameGenreRepository {
	return postgres.New(env.Postgres)
}
