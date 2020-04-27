package game

import (
	"github.com/qilin/crm-api/internal/domain/service"
	"go.uber.org/fx"
)

type Handler struct {
	GameService service.GameService
}

type Params struct {
	fx.In

	GameService service.GameService
}

func New(params Params) *Handler {
	return &Handler{
		GameService: params.GameService,
	}
}
