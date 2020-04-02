package tag

import (
	"github.com/qilin/crm-api/internal/domain/repository"
	"github.com/qilin/crm-api/internal/domain/service"
)

type ServiceParams struct {
	GameService       service.IGameService
	TagRepository     repository.TagRepository
	GameTagRepository repository.GameTagRepository
}

func New(params ServiceParams) *Service {
	return &Service{
		params,
	}
}
