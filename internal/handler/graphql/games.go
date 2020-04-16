package graphql

import (
	"context"
)

// Games returns all existing games
func (r *queryResolver) Games(ctx context.Context) ([]*Game, error) {
	// TODO +paging
	panic("not implemented")
}

// Game returns game type by id
func (r *queryResolver) Game(ctx context.Context, id string) (*Game, error) {
	g, err := r.gameService.GetExistExByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return r.convertGame(g), nil
}

// CreateGame creates new game in db
func (r *mutationResolver) CreateGame(ctx context.Context, input CreateGameInput) (*Game, error) {
	g, err := r.gameService.Create(input.convert(ctx))
	if err != nil {
		return nil, err
	}
	return r.convertGame(g), nil
}

// UpdateGame updates game data
func (r *mutationResolver) UpdateGame(ctx context.Context, input UpdateGameInput) (*Game, error) {
	g, err := r.gameService.Update(input.convert(ctx))
	if err != nil {
		return nil, err
	}
	return r.convertGame(g), nil
}

// DeleteGame removes game from database
func (r *mutationResolver) DeleteGame(ctx context.Context, id string) (bool, error) {
	err := r.gameService.Delete(ctx, id)
	return err == nil, err
}
