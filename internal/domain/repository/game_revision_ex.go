package repository

import (
	"context"

	"github.com/qilin/crm-api/internal/domain/entity"
)

type GameRevisionExRepository interface {
	FindByID(ctx context.Context, id uint) (*entity.GameRevisionEx, error)
}
