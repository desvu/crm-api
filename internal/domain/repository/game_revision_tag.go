package repository

import (
	"context"

	"github.com/qilin/crm-api/internal/domain/entity"
)

//go:generate mockgen -destination=../mocks/game_revision_tag_repository.go -package=mocks github.com/qilin/crm-api/internal/domain/repository GameRevisionTagRepository
type GameRevisionTagRepository interface {
	Create(ctx context.Context, i *entity.GameRevisionTag) error
	CreateMultiple(ctx context.Context, items []entity.GameRevisionTag) error
	Delete(ctx context.Context, i *entity.GameRevisionTag) error
	DeleteMultiple(ctx context.Context, items []entity.GameRevisionTag) error

	FindByTagID(ctx context.Context, tagID uint) ([]entity.GameRevisionTag, error)
	FindByGameRevisionID(ctx context.Context, gameRevisionID uint) ([]entity.GameRevisionTag, error)
	FindByGameRevisionIDAndTagIDs(ctx context.Context, gameRevisionID uint, tagIDs []uint) ([]entity.GameRevisionTag, error)
}
