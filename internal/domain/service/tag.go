package service

import (
	"context"

	"github.com/qilin/crm-api/internal/domain/entity"
)

//go:generate mockgen -destination=../mocks/tag_service.go -package=mocks github.com/qilin/crm-api/internal/domain/service TagService
type TagService interface {
	Create(ctx context.Context, data *CreateTagData) (*entity.Tag, error)
	Update(ctx context.Context, data *UpdateTagData) (*entity.Tag, error)
	Delete(ctx context.Context, id uint) error

	GetByID(ctx context.Context, id uint) (*entity.Tag, error)
	GetExistByID(ctx context.Context, id uint) (*entity.Tag, error)
	GetByIDs(ctx context.Context, ids []uint) ([]entity.Tag, error)
	GetByGameRevisionID(ctx context.Context, gameRevisionID uint) ([]entity.Tag, error)

	UpdateTagsForGameRevision(ctx context.Context, gameRevision *entity.GameRevision, tagIDs []uint) error
}

type CreateTagData struct {
	Name string
}

type UpdateTagData struct {
	ID   uint
	Name string
}
