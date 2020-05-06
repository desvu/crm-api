package service

import (
	"context"

	"github.com/qilin/crm-api/internal/domain/entity"
)

//go:generate mockgen -destination=../mocks/localization_service.go -package=mocks github.com/qilin/crm-api/internal/domain/service LocalizationService
type LocalizationService interface {
	GetByID(ctx context.Context, id uint) (*entity.Localization, error)
	GetExistByID(ctx context.Context, id uint) (*entity.Localization, error)
	GetByIDs(ctx context.Context, ids []uint) ([]entity.Localization, error)
	GetByGameRevisionID(ctx context.Context, gameRevisionID uint) ([]entity.Localization, error)

	UpdateLocalizationsForGameRevision(ctx context.Context, gameRevision *entity.GameRevision, localizations []LocalizationData) error
}

type LocalizationData struct {
	Language  string
	Interface bool
	Audio     bool
	Subtitles bool
}
