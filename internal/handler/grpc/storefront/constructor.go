package storefront

import (
	"github.com/qilin/crm-api/internal/domain/service"
	"go.uber.org/fx"
)

type Handler struct {
	StorefrontService service.StorefrontService
}

type Params struct {
	fx.In

	StorefrontService service.StorefrontService
}

func New(params Params) *Handler {
	return &Handler{
		StorefrontService: params.StorefrontService,
	}
}
