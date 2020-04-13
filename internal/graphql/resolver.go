package graphql

import (
	"github.com/qilin/crm-api/internal/domain/service"
	"go.uber.org/fx"
)

type Resolver struct {
	gameConverter
	games service.GameService
}

type Params struct {
	fx.In

	GameService service.GameService
}

func NewResolver(params Params) *Resolver {
	return &Resolver{games: params.GameService}
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type mutationResolver struct{ *Resolver }

type queryResolver struct{ *Resolver }
