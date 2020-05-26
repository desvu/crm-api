package repository

import (
	"context"

	"github.com/qilin/crm-api/internal/domain/entity"
)

type GameRepository interface {
	Create(ctx context.Context, i *entity.Game) error
	Update(ctx context.Context, i *entity.Game) error
	Delete(ctx context.Context, i *entity.Game) error

	FindByID(ctx context.Context, id string) (*entity.Game, error)
	FindBySlug(ctx context.Context, slug string) (*entity.Game, error)
	FindByIDs(ctx context.Context, ids []string) ([]entity.Game, error)
	FindByFilter(ctx context.Context, data *FindByFilterGameData) ([]entity.Game, error)
	FindByTitleSubstring(ctx context.Context, data *FindByTitleSubstringData) ([]entity.Game, error)
}

type FindByFilterGameData struct {
	Limit  int
	Offset int
}

type FindByTitleSubstringData struct {
	Title  string
	Limit  int
	Offset int
}
