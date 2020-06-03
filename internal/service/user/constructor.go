package user

import (
	"github.com/qilin/crm-api/internal/domain/repository"
	"github.com/qilin/crm-api/internal/domain/service"
	"go.uber.org/fx"
)

type ServiceParams struct {
	fx.In

	UserRepository         repository.UserRepository
	DocumentRepository     repository.DocumentRepository
	UserDocumentRepository repository.UserDocumentRepository
}

func New(params ServiceParams) service.UserService {
	return &Service{
		params,
	}
}
