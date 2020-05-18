package publisher

import (
	"github.com/qilin/crm-api/internal/domain/service"
	"go.uber.org/fx"
)

type Params struct {
	fx.In

	Publisher service.PublisherService
}

type Handler struct {
	Publisher service.PublisherService
}

func NewHandler(p Params) Handler {
	return Handler{p.Publisher}
}
