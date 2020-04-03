package repository

import (
	"context"

	"github.com/qilin/crm-api/internal/domain/entity"
)

//go:generate mockgen -destination=../mocks/game_developer_repository.go -package=mocks github.com/qilin/crm-api/internal/domain/repository GameDeveloperRepository
type GameDeveloperRepository interface {
	Create(ctx context.Context, i *entity.GameDeveloper) error
	CreateMultiple(ctx context.Context, items []entity.GameDeveloper) error
	Delete(ctx context.Context, i *entity.GameDeveloper) error
	DeleteMultiple(ctx context.Context, items []entity.GameDeveloper) error

	FindByDeveloperID(ctx context.Context, developerID uint) ([]entity.GameDeveloper, error)
	FindByGameID(ctx context.Context, gameID uint) ([]entity.GameDeveloper, error)
	FindByGameIDAndDeveloperIDs(ctx context.Context, gameID uint, developerIDs []uint) ([]entity.GameDeveloper, error)
}
