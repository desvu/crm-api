package document

import (
	"github.com/qilin/crm-api/internal/domain/repository"
	"github.com/qilin/crm-api/internal/domain/service"
	"go.uber.org/fx"
)

type ServiceParams struct {
	fx.In

	DocumentRepository repository.DocumentRepository
}

func New(params ServiceParams) service.DocumentService {
	return &Service{
		params,
	}
}
