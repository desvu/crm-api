package service

import (
	"context"

	"github.com/qilin/crm-api/internal/domain/entity"
)

type GameRevisionMediaService interface {
	Create(ctx context.Context, revision *entity.GameRevision, media *entity.GameMedia) (*entity.GameRevisionMedia, error)

	GetByRevision(ctx context.Context, revision *entity.GameRevision) ([]entity.GameRevisionMedia, error)
}
