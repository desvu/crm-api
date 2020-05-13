package repository

import (
	"context"

	"github.com/qilin/crm-api/internal/domain/entity"
)

//go:generate mockgen -destination=../mocks/game_revision_review_repository.go -package=mocks github.com/qilin/crm-api/internal/domain/repository GameRevisionReviewRepository
type GameRevisionReviewRepository interface {
	Create(ctx context.Context, i *entity.Review) error
	CreateMultiple(ctx context.Context, items []entity.Review) error
	Update(ctx context.Context, i *entity.Review) error
	UpdateMultiple(ctx context.Context, items []entity.Review) error
	Delete(ctx context.Context, i *entity.Review) error
	DeleteMultiple(ctx context.Context, items []entity.Review) error

	FindByID(ctx context.Context, id uint) (*entity.Review, error)
	FindByIDs(ctx context.Context, ids []uint) ([]entity.Review, error)
	FindByGameRevisionID(ctx context.Context, gameRevisionID uint) ([]entity.Review, error)
}
