package genre

import (
	"github.com/qilin/crm-api/internal/domain/repository"
	"github.com/qilin/crm-api/internal/domain/service"
	"go.uber.org/fx"
)

type ServiceParams struct {
	fx.In

	GenreRepository     repository.GenreRepository
	GameGenreRepository repository.GameGenreRepository
}

func New(params ServiceParams) service.GenreService {
	return &Service{
		params,
	}
}
