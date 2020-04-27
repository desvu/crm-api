package event

import (
	"github.com/qilin/crm-api/internal/handler/publisher/game_store"
	"go.uber.org/fx"
)

func New() fx.Option {
	return fx.Provide(
		game_store.New,
	)
}
