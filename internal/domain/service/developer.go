package service

import (
	"context"

	"github.com/qilin/crm-api/internal/domain/entity"
)

//go:generate mockgen -destination=../mocks/developer_service.go -package=mocks github.com/qilin/crm-api/internal/domain/service DeveloperService
type DeveloperService interface {
	Create(ctx context.Context, data *CreateDeveloperData) (*entity.Developer, error)
	Update(ctx context.Context, data *UpdateDeveloperData) (*entity.Developer, error)
	Delete(ctx context.Context, id uint) error

	GetByID(ctx context.Context, id uint) (*entity.Developer, error)
	GetExistByID(ctx context.Context, id uint) (*entity.Developer, error)
	GetByIDs(ctx context.Context, ids []uint) ([]entity.Developer, error)
	GetByGameRevisionID(ctx context.Context, gameRevisionID uint) ([]entity.Developer, error)

	UpdateDevelopersForGameRevision(ctx context.Context, gameRevision *entity.GameRevision, developerIDs []uint) error
}

type CreateDeveloperData struct {
	Name string
}

type UpdateDeveloperData struct {
	ID   uint
	Name string
}
