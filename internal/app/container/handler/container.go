package handler

import (
	"github.com/qilin/crm-api/internal/handler/graph"
	"github.com/qilin/crm-api/internal/handler/micro/service"
	"github.com/qilin/crm-api/internal/handler/micro/web"
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

func NewWeb() fx.Option {
	return fx.Provide(
		web.New,
	)
}
