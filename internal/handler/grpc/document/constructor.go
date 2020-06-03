package document

import (
	"github.com/qilin/crm-api/internal/domain/service"
	"go.uber.org/fx"
)

type Handler struct {
	DocumentService service.DocumentService
	UserService     service.UserService
}

type Params struct {
	fx.In

	DocumentService service.DocumentService
	UserService     service.UserService
}

func New(params Params) *Handler {
	return &Handler{
		DocumentService: params.DocumentService,
		UserService:     params.UserService,
	}
}
