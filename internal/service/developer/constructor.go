package developer

import (
	"github.com/qilin/crm-api/internal/domain/repository"
	"github.com/qilin/crm-api/internal/domain/service"
	"go.uber.org/fx"
)

type ServiceParams struct {
	fx.In

	DeveloperRepository     repository.DeveloperRepository
	GameDeveloperRepository repository.GameDeveloperRepository
}

func New(params ServiceParams) service.DeveloperService {
	return &Service{
		params,
	}
}
