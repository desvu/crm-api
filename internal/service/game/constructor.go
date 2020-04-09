package game

import (
	"github.com/qilin/crm-api/internal/domain/repository"
	"github.com/qilin/crm-api/internal/domain/service"
	"github.com/qilin/crm-api/pkg/transactor"
)

type ServiceParams struct {
	TagService       service.TagService
	DeveloperService service.DeveloperService
	PublisherService service.PublisherService
	FeatureService   service.FeatureService
	GenreService     service.GenreService
	GameRepository   repository.GameRepository
	GameExRepository repository.GameExRepository
	Transactor       *transactor.Transactor
}

func New(params ServiceParams) *Service {
	return &Service{
		params,
	}
}
