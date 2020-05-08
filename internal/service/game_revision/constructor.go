package game_revision

import (
	"github.com/qilin/crm-api/internal/domain/repository"
	"github.com/qilin/crm-api/internal/domain/service"
	"github.com/qilin/crm-api/pkg/transactor"
	"go.uber.org/fx"
)

type ServiceParams struct {
	fx.In

	TagService               service.TagService
	DeveloperService         service.DeveloperService
	PublisherService         service.PublisherService
	FeatureService           service.FeatureService
	GenreService             service.GenreService
	GameMediaService         service.GameMediaService
	GameRevisionRepository   repository.GameRevisionRepository
	GameRevisionExRepository repository.GameRevisionExRepository
	Transactor               *transactor.Transactor
}

func New(params ServiceParams) service.GameRevisionService {
	return &Service{
		params,
	}
}
