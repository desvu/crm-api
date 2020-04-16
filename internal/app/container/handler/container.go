package handler

import (
	"github.com/qilin/crm-api/internal/handler/graph"
	"go.uber.org/fx"
)

func New() fx.Option {
	return fx.Provide(
		graph.NewResolver,
	)
}
