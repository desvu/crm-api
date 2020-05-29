package document

import (
	"github.com/qilin/crm-api/internal/domain/service"
	"go.uber.org/fx"
)

type Params struct {
	fx.In

	DocumentService service.DocumentService
}

type Handler struct {
	DocumentService service.DocumentService
}

func NewHandler(params Params) Handler {
	return Handler{
		DocumentService: params.DocumentService,
	}
}
