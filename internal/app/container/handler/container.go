package handler

import (
	"github.com/qilin/crm-api/internal/handler/graph"
	"github.com/qilin/crm-api/internal/handler/http"
	"github.com/qilin/crm-api/internal/handler/micro/service"
	"go.uber.org/fx"
)

func New() fx.Option {
	return fx.Provide(
		graph.NewResolver,
	)
}

func NewGRPC() fx.Option {
	return fx.Provide(
		service.New,
	)
}

func NewHTTP() fx.Option {
	return fx.Provide(
		http.New,
	)
}
