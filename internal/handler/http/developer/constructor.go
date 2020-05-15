package developer

import (
	"github.com/qilin/crm-api/internal/domain/service"
	"go.uber.org/fx"
)

type Params struct {
	fx.In

	Developer service.DeveloperService
}

type Handler struct {
	Developer service.DeveloperService
}

func NewHandler(p Params) Handler {
	return Handler{p.Developer}
}
