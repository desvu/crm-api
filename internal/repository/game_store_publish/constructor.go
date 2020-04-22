package game_store_publish

import (
	"github.com/qilin/crm-api/internal/domain/repository"
	"github.com/qilin/crm-api/internal/env"
	"github.com/qilin/crm-api/internal/repository/game_store_publish/postgres"
)

func New(env *env.Env) repository.GameStorePublishRepository {
	return postgres.New(env.Store.Postgres)
}
