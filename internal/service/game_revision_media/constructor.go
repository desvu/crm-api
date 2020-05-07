package game_revision_media

import (
	"github.com/qilin/crm-api/internal/domain/repository"
	"github.com/qilin/crm-api/internal/domain/service"
	"go.uber.org/fx"
)

type ServiceParams struct {
	fx.In

	GameRevisionMediaRepository repository.GameRevisionMediaRepository
}

func New(params ServiceParams) service.GameRevisionMediaService {
	return &Service{
		params,
	}
}
