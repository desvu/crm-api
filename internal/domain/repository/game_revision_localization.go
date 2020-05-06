package repository

import (
	"context"

	"github.com/qilin/crm-api/internal/domain/entity"
)

//go:generate mockgen -destination=../mocks/game_revision_localization_repository.go -package=mocks github.com/qilin/crm-api/internal/domain/repository GameRevisionLocalizationRepository
type GameRevisionLocalizationRepository interface {
	Create(ctx context.Context, i *entity.Localization) error
	CreateMultiple(ctx context.Context, items []entity.Localization) error
	Update(ctx context.Context, i *entity.Localization) error
	UpdateMultiple(ctx context.Context, items []entity.Localization) error
	Delete(ctx context.Context, i *entity.Localization) error
	DeleteMultiple(ctx context.Context, items []entity.Localization) error

	FindByID(ctx context.Context, id uint) (*entity.Localization, error)
	FindByIDs(ctx context.Context, ids []uint) ([]entity.Localization, error)
	FindByGameRevisionID(ctx context.Context, gameRevisionID uint) ([]entity.Localization, error)
	FindByGameRevisionIDAndLanguage(ctx context.Context, gameRevisionID uint, langs []string) ([]entity.Localization, error)
}
