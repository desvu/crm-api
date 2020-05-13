package review

import (
	"github.com/qilin/crm-api/internal/domain/repository"
	"github.com/qilin/crm-api/internal/domain/service"
	"go.uber.org/fx"
)

type ServiceParams struct {
	fx.In

	ReviewRepository repository.GameRevisionReviewRepository
}

func New(params ServiceParams) service.ReviewService {
	return &Service{
		params,
	}
}
