package game_publisher

import (
	"github.com/qilin/crm-api/internal/domain/repository"
	"github.com/qilin/crm-api/internal/env"
	"github.com/qilin/crm-api/internal/repository/game_publisher/postgres"
)

func New(env *env.Store) repository.GamePublisherRepository {
	return postgres.New(env.Postgres)
}
