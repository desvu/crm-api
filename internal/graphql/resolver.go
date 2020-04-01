package graphql

import (
	"github.com/qilin/crm-api/internal/domain/service"
)

type Resolver struct {
	games service.IGameService
}

func NewResolver(games service.IGameService) *Resolver {
	return &Resolver{games: games}
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
