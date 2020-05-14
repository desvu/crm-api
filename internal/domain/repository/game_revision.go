package repository

import (
	"context"

	"github.com/qilin/crm-api/internal/domain/entity"
)

type GameRevisionRepository interface {
	Create(ctx context.Context, i *entity.GameRevision) error
	Update(ctx context.Context, i *entity.GameRevision) error
	Delete(ctx context.Context, i *entity.GameRevision) error

	FindByID(ctx context.Context, id uint) (*entity.GameRevision, error)
	FindByIDs(ctx context.Context, ids []uint) ([]entity.GameRevision, error)
	FindByGameID(ctx context.Context, gameID string) ([]entity.GameRevision, error)
	FindLastByGameIDs(ctx context.Context, gameIDs []string) ([]entity.GameRevision, error)
	FindByIDAndGameID(ctx context.Context, id uint, gameID string) (*entity.GameRevision, error)
	FindDraftByGameID(ctx context.Context, gameID string) (*entity.GameRevision, error)
	FindLastPublishedByGameID(ctx context.Context, gameID string) (*entity.GameRevision, error)
	FindPublishedByGameIDs(ctx context.Context, gameIDs []string) ([]string, error)
}
