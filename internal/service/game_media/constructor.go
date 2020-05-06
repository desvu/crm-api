package game_media

import (
	"github.com/qilin/crm-api/internal/domain/repository"
	"github.com/qilin/crm-api/internal/domain/service"
	"github.com/qilin/crm-api/internal/env/storage"
	"go.uber.org/fx"
)

type ServiceParams struct {
	fx.In

	GameService                 service.GameService
	GameMediaRepository         repository.GameMediaRepository
	GameRevisionMediaRepository repository.GameRevisionMediaRepository
	Storage                     *storage.Env
}

func New(params ServiceParams) service.GameMediaService {
	return &Service{
		params,
	}
}
