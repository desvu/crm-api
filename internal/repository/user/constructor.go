package user

import (
	"github.com/qilin/crm-api/internal/domain/repository"
	"github.com/qilin/crm-api/internal/env"
	"github.com/qilin/crm-api/internal/repository/user/postgres"
)

func New(env *env.Env) repository.UserRepository {
	return postgres.New(env.Store.Postgres)
}
