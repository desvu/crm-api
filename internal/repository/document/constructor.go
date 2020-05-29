package document

import (
	"github.com/qilin/crm-api/internal/domain/repository"
	"github.com/qilin/crm-api/internal/env"
	"github.com/qilin/crm-api/internal/repository/document/postgres"
)

func New(env *env.Env) repository.DocumentRepository {
	return postgres.New(env.Store.Postgres)
}
