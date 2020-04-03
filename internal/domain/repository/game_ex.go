package repository

import (
	"context"

	"github.com/qilin/crm-api/internal/domain/entity"
)

type GameExRepository interface {
	FindByID(ctx context.Context, id uint) (*entity.GameEx, error)
}
