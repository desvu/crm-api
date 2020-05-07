package storefront

import (
	"github.com/qilin/crm-api/internal/domain/repository"
	"github.com/qilin/crm-api/internal/domain/service"
	"github.com/qilin/crm-api/pkg/transactor"
	"go.uber.org/fx"
)

type ServiceParams struct {
	fx.In

	GameRevisionService service.GameRevisionService
	Repository          repository.StoreFrontRepository
	Transactor          *transactor.Transactor
}

func New(params ServiceParams) service.StorefrontService {
	return &Service{
		params,
	}
}
