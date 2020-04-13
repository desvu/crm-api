package tag

import (
	"github.com/qilin/crm-api/internal/domain/repository"
	"github.com/qilin/crm-api/internal/domain/service"
	"go.uber.org/fx"
)

type ServiceParams struct {
	fx.In

	TagRepository     repository.TagRepository
	GameTagRepository repository.GameTagRepository
}

func New(params ServiceParams) service.TagService {
	return &Service{
		params,
	}
}
