package micro

import (
	"github.com/micro/go-micro/web"
	"github.com/qilin/crm-api/internal/handler/graph"
	"go.uber.org/fx"
)

type Params struct {
	fx.In

	Handler *graph.Resolver
}

func New(params Params) (web.Service, error) {
	s := web.NewService(
		web.Name("p1.crm.api"),
		web.Version("latest"),
		web.Address(":8080"),
	)

	s.Handle("/api/graphql/client", graph.Playground("/api/graphql"))
	s.Handle("/api/graphql", graph.NewHandler(params.Handler))

	if err := s.Init(); err != nil {
		return nil, err
	}

	return s, nil
}
