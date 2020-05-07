package handler

import (
	"github.com/qilin/crm-api/internal/handler/graph"
	"github.com/qilin/crm-api/internal/handler/grpc"
	"github.com/qilin/crm-api/internal/handler/grpc/game"
	"github.com/qilin/crm-api/internal/handler/http"
	"github.com/qilin/crm-api/internal/handler/http/storefront"
	"go.uber.org/fx"
)

func New() fx.Option {
	return fx.Provide(
		storefront.NewHandler,
		graph.NewResolver,
		game.New,
		grpc.New,
		http.New,
	)
}
