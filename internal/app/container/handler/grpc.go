package handler

import (
	"github.com/qilin/crm-api/internal/handler/grpc"
	"github.com/qilin/crm-api/internal/handler/grpc/document"
	"github.com/qilin/crm-api/internal/handler/grpc/game"
	"github.com/qilin/crm-api/internal/handler/grpc/storefront"
	"go.uber.org/fx"
)

func NewGrpc() fx.Option {
	return fx.Provide(
		grpc.New,
		game.New,
		storefront.New,
		document.New,
	)
}
