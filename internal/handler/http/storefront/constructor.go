package storefront

import (
	"github.com/qilin/crm-api/internal/domain/service"
	"go.uber.org/fx"
)

type Params struct {
	fx.In

	Storefronts service.StorefrontService
}

type Handler struct {
	Storefronts service.StorefrontService
}

func NewHandler(p Params) *Handler {
	return &Handler{p.Storefronts}
}
