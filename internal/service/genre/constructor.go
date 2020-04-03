package genre

import (
	"github.com/qilin/crm-api/internal/domain/repository"
	"github.com/qilin/crm-api/internal/domain/service"
)

type ServiceParams struct {
	GameService         service.GameService
	GenreRepository     repository.GenreRepository
	GameGenreRepository repository.GameGenreRepository
}

func New(params ServiceParams) *Service {
	return &Service{
		params,
	}
}
