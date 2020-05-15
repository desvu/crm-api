package handler

import (
	"github.com/qilin/crm-api/internal/handler/http"
	"github.com/qilin/crm-api/internal/handler/http/feature"
	"github.com/qilin/crm-api/internal/handler/http/genre"
	"github.com/qilin/crm-api/internal/handler/http/storefront"
	"github.com/qilin/crm-api/internal/handler/http/tag"
	"go.uber.org/fx"
)

func NewHttp() fx.Option {
	return fx.Provide(
		http.New,
		storefront.NewHandler,
		tag.NewHandler,
		genre.NewHandler,
		feature.NewHandler,
	)
}
