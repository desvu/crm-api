package rating

import (
	"github.com/qilin/crm-api/internal/domain/repository"
	"github.com/qilin/crm-api/internal/domain/service"
	"go.uber.org/fx"
)

type ServiceParams struct {
	fx.In

	RatingRepository repository.GameRevisionRatingRepository
}

func New(params ServiceParams) service.RatingService {
	return &Service{
		params,
	}
}
