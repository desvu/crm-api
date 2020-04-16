package graph

import (
	"github.com/qilin/crm-api/internal/domain/service"
	"github.com/qilin/crm-api/internal/handler/graph/generated"
	"go.uber.org/fx"
)

type Resolver struct {
	gameConverter
	gameService service.GameService
}

type Params struct {
	fx.In

	GameService service.GameService
}

func NewResolver(params Params) *Resolver {
	return &Resolver{gameService: params.GameService}
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver {
	return &mutationResolver{r}
}

// Query returns QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver {
	return &queryResolver{r}
}

type mutationResolver struct{ *Resolver }

type queryResolver struct{ *Resolver }
