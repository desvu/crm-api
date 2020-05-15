package handler

import (
	"github.com/qilin/crm-api/internal/handler/http"
	"github.com/qilin/crm-api/internal/handler/http/developer"
	"github.com/qilin/crm-api/internal/handler/http/storefront"
	"go.uber.org/fx"
)

func NewHttp() fx.Option {
	return fx.Provide(
		http.New,
		storefront.NewHandler,
		developer.NewHandler,
	)
}
