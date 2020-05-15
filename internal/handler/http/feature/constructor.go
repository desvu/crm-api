package feature

import (
	"github.com/qilin/crm-api/internal/domain/service"
	"go.uber.org/fx"
)

type Params struct {
	fx.In

	Feature service.FeatureService
}

type Handler struct {
	Feature service.FeatureService
}

func NewHandler(p Params) Handler {
	return Handler{p.Feature}
}
