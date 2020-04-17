package game

import (
	"github.com/qilin/crm-api/internal/domain/repository"
	"github.com/qilin/crm-api/internal/domain/service"
	"github.com/qilin/crm-api/pkg/transactor"
	"go.uber.org/fx"
)

type ServiceParams struct {
	fx.In

	GameRevisionService service.GameRevisionService
	TagService          service.TagService
	DeveloperService    service.DeveloperService
	PublisherService    service.PublisherService
	FeatureService      service.FeatureService
	GenreService        service.GenreService
	GameRepository      repository.GameRepository
	GameExRepository    repository.GameRevisionExRepository
	Transactor          *transactor.Transactor
}

func New(params ServiceParams) service.GameService {
	return &Service{
		params,
	}
}
