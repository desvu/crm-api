package user_document

import (
	"github.com/qilin/crm-api/internal/domain/repository"
	"github.com/qilin/crm-api/internal/env"
	"github.com/qilin/crm-api/internal/repository/user_document/postgres"
)

func New(env *env.Env) repository.UserDocumentRepository {
	return postgres.New(env.Store.Postgres)
}
