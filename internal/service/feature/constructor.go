package feature

import (
	"github.com/qilin/crm-api/internal/domain/repository"
	"github.com/qilin/crm-api/internal/domain/service"
)

type ServiceParams struct {
	GameService           service.GameService
	FeatureRepository     repository.FeatureRepository
	GameFeatureRepository repository.GameFeatureRepository
}

func New(params ServiceParams) *Service {
	return &Service{
		params,
	}
}
