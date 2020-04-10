package handler

import (
	"github.com/qilin/crm-api/internal/graphql"
	"go.uber.org/fx"
)

func New() fx.Option {
	return fx.Provide(
		graphql.NewResolver,
	)
}
