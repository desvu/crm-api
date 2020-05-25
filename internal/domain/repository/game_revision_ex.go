package repository

import (
	"context"

	"github.com/qilin/crm-api/internal/domain/entity"
)

type GameRevisionExRepository interface {
	FindByID(ctx context.Context, id uint) (*entity.GameRevisionEx, error)
	FindLastByGameIDs(ctx context.Context, gameIDs []string) ([]entity.GameRevisionEx, error)
	FindByIDAndGameID(ctx context.Context, id uint, gameID string) (*entity.GameRevisionEx, error)
	FindLastPublishedByGameID(ctx context.Context, gameID string) (*entity.GameRevisionEx, error)
}
