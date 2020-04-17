package game_revision_publisher

import (
	"github.com/qilin/crm-api/internal/domain/repository"
	"github.com/qilin/crm-api/internal/env"
	"github.com/qilin/crm-api/internal/repository/game_revision_publisher/postgres"
)

func New(env *env.Env) repository.GameRevisionPublisherRepository {
	return postgres.New(env.Store.Postgres)
}
