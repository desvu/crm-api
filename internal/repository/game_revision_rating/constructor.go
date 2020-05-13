package game_revision_rating

import (
	"github.com/qilin/crm-api/internal/domain/repository"
	"github.com/qilin/crm-api/internal/env"
	"github.com/qilin/crm-api/internal/repository/game_revision_rating/postgres"
)

func New(env *env.Env) repository.GameRevisionRatingRepository {
	return postgres.New(env.Store.Postgres)
}
