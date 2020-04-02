package graphql

import (
	"context"
	"encoding/json"
	"strconv"
	"time"

	"github.com/qilin/crm-api/internal/domain/entity"
	"github.com/qilin/crm-api/internal/domain/enum/game"
	"github.com/qilin/crm-api/internal/domain/service"
)

// Games returns all existing games
func (r *queryResolver) Games(ctx context.Context) ([]*Game, error) {
	// TODO +paging
	panic("not implemented")
}

// Game returns game type by id
func (r *queryResolver) Game(ctx context.Context, id string) (*Game, error) {
	return r.getGame(ctx, id)
}

// CreateGame creates new game in db
func (r *mutationResolver) CreateGame(ctx context.Context, title string) (*Game, error) {
	g, err := r.games.Create(ctx, &service.CreateGameData{
		Title:       title,
		Summary:     "xxx",
		Description: "zzz",
		License:     "qqq",
		Ranking:     "rrr",
		Type:        game.TypeDesktop,
		Platforms:   []game.Platform{game.PlatformWindows, game.PlatformWeb},
		ReleaseDate: time.Now(),
	})
	if err != nil {
		return nil, err
	}
	return r.convertGame(g), nil
}

// UpdateGame updates game data
func (r *mutationResolver) UpdateGame(ctx context.Context, id string, gameJSON string) (*Game, error) {
	var g Game
	if err := json.Unmarshal([]byte(gameJSON), &g); err != nil {
		return nil, err
	}
	gid, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return nil, err
	}
	r.games.Update(ctx, &service.UpdateGameData{
		ID:    uint(gid),
		Title: &g.Title,
		// TODO more fields
	})
	return r.getGame(ctx, id)
}

// DeleteGame removes game from database
func (r *mutationResolver) DeleteGame(ctx context.Context, id string) (bool, error) {
	gid, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return false, err
	}
	err = r.games.Delete(ctx, uint(gid))
	return err == nil, err
}

func (r *Resolver) getGame(ctx context.Context, id string) (*Game, error) {
	gid, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return nil, err
	}
	g, err := r.games.GetByID(ctx, uint(gid))
	if err != nil {
		return nil, err
	}
	return r.convertGame(g), nil
}

func (r *Resolver) convertGame(g *entity.Game) *Game {
	return &Game{
		ID:    strconv.FormatUint(uint64(g.ID), 10),
		Title: g.Title,
		// TODO more fields
	}
}
