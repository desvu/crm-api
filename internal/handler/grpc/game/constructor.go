package game

import (
	"github.com/qilin/crm-api/internal/domain/service"
	"github.com/qilin/crm-api/internal/helper/url_builder"
	"go.uber.org/fx"
)

type Handler struct {
	GameService service.GameService
	URLBuilder  url_builder.Helper
}

type Params struct {
	fx.In

	GameService service.GameService
	URLBuilder  url_builder.Helper
}

func New(params Params) *Handler {
	return &Handler{
		GameService: params.GameService,
		URLBuilder:  params.URLBuilder,
	}
}
