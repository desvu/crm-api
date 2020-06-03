package service

import (
	"context"

	"github.com/qilin/crm-api/internal/domain/entity"
)

type UserService interface {
	GetByExternalID(ctx context.Context, externalID string) (*entity.User, error)
}
