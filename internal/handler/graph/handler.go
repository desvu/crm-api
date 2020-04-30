package graph

import (
	"context"
	"net/http"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/handler"
	"github.com/qilin/crm-api/internal/handler/graph/generated"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"go.uber.org/zap"
)

func Playground(endpoint string) http.Handler {
	return handler.Playground("local", endpoint)
}

func NewHandler(resolver *Resolver) http.Handler {
	options := []handler.Option{
		handler.IntrospectionEnabled(true),
		handler.ErrorPresenter(func(ctx context.Context, err error) *gqlerror.Error {
			zap.L().Error("GraphQL query failed", zap.Error(err))
			return graphql.DefaultErrorPresenter(ctx, err)
		}),
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
