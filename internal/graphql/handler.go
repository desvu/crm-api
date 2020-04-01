package graphql

import (
	"net/http"

	"github.com/99designs/gqlgen/handler"
)

func NewHandler(resolver *Resolver) http.HandlerFunc {
	options := []handler.Option{}

	cfg := Config{
		Resolvers: resolver,
	}

	h := handler.GraphQL(
		NewExecutableSchema(cfg),
		options...,
	)

	return h
}
