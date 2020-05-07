package handler

import (
	"github.com/qilin/crm-api/internal/handler/grpc"
	"github.com/qilin/crm-api/internal/handler/grpc/game"
	"github.com/qilin/crm-api/internal/handler/http"
	"go.uber.org/fx"
)

func New() fx.Option {
	return fx.Provide(
		game.New,
		grpc.New,
		http.New,
	)
}
