package repository

import (
	"context"

	"github.com/qilin/crm-api/internal/domain/entity"
)

//go:generate mockgen -destination=../mocks/game_revision_rating_repository.go -package=mocks github.com/qilin/crm-api/internal/domain/repository GameRevisionRatingRepository
type GameRevisionRatingRepository interface {
	Create(ctx context.Context, i *entity.Rating) error
	CreateMultiple(ctx context.Context, items []entity.Rating) error
	Update(ctx context.Context, i *entity.Rating) error
	UpdateMultiple(ctx context.Context, items []entity.Rating) error
	Delete(ctx context.Context, i *entity.Rating) error
	DeleteMultiple(ctx context.Context, items []entity.Rating) error

	FindByID(ctx context.Context, id uint) (*entity.Rating, error)
	FindByIDs(ctx context.Context, ids []uint) ([]entity.Rating, error)
	FindByGameRevisionID(ctx context.Context, gameRevisionID uint) ([]entity.Rating, error)
	FindByGameRevisionIDAndAgency(ctx context.Context, gameRevisionID uint, agencies []uint8) ([]entity.Rating, error)
}
