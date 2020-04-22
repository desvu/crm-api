package game_store_publish

import (
	"github.com/qilin/crm-api/internal/domain/repository"
	"github.com/qilin/crm-api/internal/domain/service"
	"go.uber.org/fx"
)

type ServiceParams struct {
	fx.In

	GameStorePublishRepository repository.GameStorePublishRepository
}

func New(params ServiceParams) service.GameStorePublishService {
	return &Service{
		params,
	}
}
