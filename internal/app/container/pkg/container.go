package pkg

import (
	"github.com/qilin/crm-api/pkg/broker"
	"github.com/qilin/crm-api/pkg/transactor"
	"go.uber.org/fx"
)

func New() fx.Option {
	return fx.Provide(
		transactor.New,
		broker.New,
	)
}
