package game_revision_media

import (
	"context"

	"github.com/qilin/crm-api/internal/domain/entity"
)

type Service struct {
	ServiceParams
}

func (s Service) Create(ctx context.Context, revision *entity.GameRevision, media *entity.GameMedia) (*entity.GameRevisionMedia, error) {
	grm := &entity.GameRevisionMedia{
		RevisionID: revision.ID,
		MediaID:    media.ID,
	}

	if err := s.GameRevisionMediaRepository.Create(ctx, grm); err != nil {
		return nil, err
	}

	return grm, nil
}

func (s Service) GetByRevision(ctx context.Context, revision *entity.GameRevision) ([]entity.GameRevisionMedia, error) {
	return s.GameRevisionMediaRepository.FindByRevisionID(ctx, revision.ID)
}
