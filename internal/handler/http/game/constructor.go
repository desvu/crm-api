package game

import (
	"github.com/qilin/crm-api/internal/domain/service"
	"github.com/qilin/crm-api/internal/helper/url_builder"
	"go.uber.org/fx"
)

type Params struct {
	fx.In

	GameService service.GameService
	URLBuilder  url_builder.Helper
}

type Handler struct {
	Params
}

func New(params Params) Handler {
	return Handler{
		params,
	}
}
