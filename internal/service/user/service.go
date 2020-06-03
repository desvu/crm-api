package user

import (
	"context"

	"github.com/qilin/crm-api/internal/domain/entity"
)

type Service struct {
	ServiceParams
}

func (s *Service) GetByExternalID(ctx context.Context, externalID string) (*entity.User, error) {
	return s.UserRepository.FindByExternalID(ctx, externalID)
}
