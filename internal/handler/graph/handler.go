package graph

import (
	"net/http"

	"github.com/99designs/gqlgen/handler"
	"github.com/qilin/crm-api/internal/handler/graph/generated"
)

func Playground(endpoint string) http.Handler {
	return handler.Playground("local", endpoint)
}

func NewHandler(resolver *Resolver) http.Handler {
	options := []handler.Option{
		handler.IntrospectionEnabled(true),
	}

	cfg := generated.Config{
		Resolvers: resolver,
	}

	h := handler.GraphQL(
		generated.NewExecutableSchema(cfg),
		options...,
	)

	return h
}
