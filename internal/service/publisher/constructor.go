package publisher

import (
	"github.com/qilin/crm-api/internal/domain/repository"
	"github.com/qilin/crm-api/internal/domain/service"
	"go.uber.org/fx"
)

type ServiceParams struct {
	fx.In

	PublisherRepository             repository.PublisherRepository
	GameRevisionPublisherRepository repository.GameRevisionPublisherRepository
}

func New(params ServiceParams) service.PublisherService {
	return &Service{
		params,
	}
}
