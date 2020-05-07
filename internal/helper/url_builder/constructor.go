package url_builder

import (
	"github.com/qilin/crm-api/internal/env"
	"go.uber.org/fx"
)

type Params struct {
	fx.In

	Env *env.Env
}

type Helper struct {
	Params
}

func New(params Params) Helper {
	return Helper{
		params,
	}
}
