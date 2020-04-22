package web

import (
	"github.com/micro/go-micro/v2/web"
	"github.com/qilin/crm-api/internal/handler/graph"
	"go.uber.org/fx"
)

type Params struct {
	fx.In

	Handler *graph.Resolver
}

func New(params Params) (web.Service, error) {
	server := web.NewService(
		web.Name("qilin.crm.web"),
		web.Version("latest"),
		web.Address(":8080"),
	)

	server.Handle("/api/graphql/client", graph.Playground("/api/graphql"))
	server.Handle("/api/graphql", graph.NewHandler(params.Handler))

	return server, nil
}
