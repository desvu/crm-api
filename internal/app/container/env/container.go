package env

import (
	"github.com/qilin/crm-api/internal/env"
	"go.uber.org/fx"
)

func New() fx.Option {
	return fx.Provide(
		env.New,
	)
}
