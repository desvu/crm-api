package tag

import (
	"github.com/qilin/crm-api/internal/domain/service"
	"go.uber.org/fx"
)

type Params struct {
	fx.In

	Tag service.TagService
}

type Handler struct {
	Tag service.TagService
}

func NewHandler(p Params) Handler {
	return Handler{p.Tag}
}
