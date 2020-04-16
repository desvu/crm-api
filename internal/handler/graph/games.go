package graph

import (
	"context"

	"github.com/qilin/crm-api/internal/handler/graph/model"
)

// Games returns all existing games
func (r *queryResolver) Games(ctx context.Context) ([]*model.Game, error) {
	// TODO +paging
	panic("not implemented")
}

// Game returns game type by id
func (r *queryResolver) Game(ctx context.Context, id string) (*model.Game, error) {
	g, err := r.gameService.GetExistExByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return r.convertGame(g), nil
}

// CreateGame creates new game in db
func (r *mutationResolver) CreateGame(ctx context.Context, input model.CreateGameInput) (*model.Game, error) {
	data, err := input.Convert()
	if err != nil {
		return nil, err
	}

	g, err := r.gameService.Create(ctx, data)
	if err != nil {
		return nil, err
	}

	return r.convertGame(g), nil
}

// UpdateGame updates game data
func (r *mutationResolver) UpdateGame(ctx context.Context, input model.UpdateGameInput) (*model.Game, error) {
	data, err := input.Convert()
	if err != nil {
		return nil, err
	}

	g, err := r.gameService.Update(ctx, data)
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
