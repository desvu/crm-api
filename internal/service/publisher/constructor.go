package publisher

import (
	"github.com/qilin/crm-api/internal/domain/repository"
	"github.com/qilin/crm-api/internal/domain/service"
)

type ServiceParams struct {
	GameService             service.GameService
	PublisherRepository     repository.PublisherRepository
	GamePublisherRepository repository.GamePublisherRepository
}

func New(params ServiceParams) *Service {
	return &Service{
		params,
	}
}
