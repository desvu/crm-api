package genre

import (
	"github.com/qilin/crm-api/internal/domain/service"
	"go.uber.org/fx"
)

type Params struct {
	fx.In

	Genre service.GenreService
}

type Handler struct {
	Genre service.GenreService
}

func NewHandler(p Params) Handler {
	return Handler{p.Genre}
}
