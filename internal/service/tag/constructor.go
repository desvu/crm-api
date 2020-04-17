package tag

import (
	"github.com/qilin/crm-api/internal/domain/repository"
	"github.com/qilin/crm-api/internal/domain/service"
	"go.uber.org/fx"
)

type ServiceParams struct {
	fx.In

	TagRepository             repository.TagRepository
	GameRevisionTagRepository repository.GameRevisionTagRepository
}

func New(params ServiceParams) service.TagService {
	return &Service{
		params,
	}
}
