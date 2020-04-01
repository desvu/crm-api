package graphql

import (
	"net/http"

	"github.com/99designs/gqlgen/handler"
)

func Playground(endpoint string) http.Handler {
	return handler.Playground("local", endpoint)
}

func NewHandler(resolver *Resolver) http.Handler {
	options := []handler.Option{
		handler.IntrospectionEnabled(true),
	}

	cfg := Config{
		Resolvers: resolver,
	}

	h := handler.GraphQL(
		NewExecutableSchema(cfg),
		options...,
	)

	return h
}
