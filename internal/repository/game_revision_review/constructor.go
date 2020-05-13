package game_revision_review

import (
	"github.com/qilin/crm-api/internal/domain/repository"
	"github.com/qilin/crm-api/internal/env"
	"github.com/qilin/crm-api/internal/repository/game_revision_review/postgres"
)

func New(env *env.Env) repository.GameRevisionReviewRepository {
	return postgres.New(env.Store.Postgres)
}
