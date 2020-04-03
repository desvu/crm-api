package developer

import (
	"github.com/qilin/crm-api/internal/domain/repository"
	"github.com/qilin/crm-api/internal/domain/service"
)

type ServiceParams struct {
	GameService             service.GameService
	DeveloperRepository     repository.DeveloperRepository
	GameDeveloperRepository repository.GameDeveloperRepository
}

func New(params ServiceParams) *Service {
	return &Service{
		params,
	}
}
