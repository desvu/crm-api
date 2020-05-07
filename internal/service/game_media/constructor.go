package game_media

import (
	"github.com/qilin/crm-api/internal/domain/repository"
	"github.com/qilin/crm-api/internal/domain/service"
	"github.com/qilin/crm-api/internal/env"
	"go.uber.org/fx"
)

type ServiceParams struct {
	fx.In

	GameService              service.GameService
	GameRevisionMediaService service.GameRevisionMediaService
	GameMediaRepository      repository.GameMediaRepository
	Env                      *env.Env
}

func New(params ServiceParams) service.GameMediaService {
	return &Service{
		params,
	}
}
