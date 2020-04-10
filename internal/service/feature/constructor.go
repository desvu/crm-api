package feature

import (
	"github.com/qilin/crm-api/internal/domain/repository"
	"github.com/qilin/crm-api/internal/domain/service"
	"go.uber.org/fx"
)

type ServiceParams struct {
	fx.In

	FeatureRepository     repository.FeatureRepository
	GameFeatureRepository repository.GameFeatureRepository
}

func New(params ServiceParams) service.FeatureService {
	return &Service{
		params,
	}
}
