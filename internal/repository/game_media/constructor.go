package game_media

import (
	"github.com/qilin/crm-api/internal/domain/repository"
	"github.com/qilin/crm-api/internal/env"
	"github.com/qilin/crm-api/internal/repository/game_media/postgres"
)

func New(env *env.Env) repository.GameMediaRepository {
	return postgres.New(env.Store.Postgres)
}
