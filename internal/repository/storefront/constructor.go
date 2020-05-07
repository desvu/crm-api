package storefront

import (
	"github.com/qilin/crm-api/internal/domain/repository"
	"github.com/qilin/crm-api/internal/env"
	"github.com/qilin/crm-api/internal/repository/storefront/postgres"
	"github.com/qilin/crm-api/pkg/transactor"
)

func New(env *env.Env, tx *transactor.Transactor) repository.StoreFrontRepository {
	return postgres.New(env.Store.Postgres, tx)
}
