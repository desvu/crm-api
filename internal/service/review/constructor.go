package review

import (
	"github.com/qilin/crm-api/internal/domain/repository"
	"github.com/qilin/crm-api/internal/domain/service"
	"github.com/qilin/crm-api/pkg/transactor"
	"go.uber.org/fx"
)

type ServiceParams struct {
	fx.In

	ReviewRepository repository.GameRevisionReviewRepository
	Transactor       *transactor.Transactor
}

func New(params ServiceParams) service.ReviewService {
	return &Service{
		params,
	}
}
