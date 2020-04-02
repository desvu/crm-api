package repository

import (
	"context"

	"github.com/qilin/crm-api/internal/domain/entity"
)

//go:generate mockgen -destination=../mocks/game_tag_repository.go -package=mocks github.com/qilin/crm-api/internal/domain/repository IGameTagRepository
type IGameTagRepository interface {
	Create(ctx context.Context, i *entity.GameTag) error
	CreateMultiple(ctx context.Context, items []entity.GameTag) error
	Delete(ctx context.Context, i *entity.GameTag) error
	DeleteMultiple(ctx context.Context, items []entity.GameTag) error

	FindByTagID(ctx context.Context, tagID uint) ([]entity.GameTag, error)
	FindByGameID(ctx context.Context, gameID uint) ([]entity.GameTag, error)
	FindByGameIDAndTagIDs(ctx context.Context, gameID uint, tagIDs []uint) ([]entity.GameTag, error)
}
