package repository

import (
	"context"

	"github.com/qilin/crm-api/internal/domain/entity"
)

//go:generate mockgen -destination=../mocks/game_revision_developer_repository.go -package=mocks github.com/qilin/crm-api/internal/domain/repository GameRevisionDeveloperRepository
type GameRevisionDeveloperRepository interface {
	Create(ctx context.Context, i *entity.GameRevisionDeveloper) error
	CreateMultiple(ctx context.Context, items []entity.GameRevisionDeveloper) error
	Delete(ctx context.Context, i *entity.GameRevisionDeveloper) error
	DeleteMultiple(ctx context.Context, items []entity.GameRevisionDeveloper) error

	FindByDeveloperID(ctx context.Context, developerID uint) ([]entity.GameRevisionDeveloper, error)
	FindByGameRevisionID(ctx context.Context, gameRevisionID uint) ([]entity.GameRevisionDeveloper, error)
	FindByGameRevisionIDs(ctx context.Context, gameRevisionIDs []uint) ([]entity.GameRevisionDeveloper, error)
	FindByGameRevisionIDAndDeveloperIDs(ctx context.Context, gameRevisionID uint, developerIDs []uint) ([]entity.GameRevisionDeveloper, error)
}
