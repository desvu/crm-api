package service

import (
	"context"

	"github.com/qilin/crm-api/internal/domain/entity"
	"github.com/qilin/crm-api/internal/domain/errors"
	"golang.org/x/text/language"
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

func (d LocalizationData) Validate() error {
	// check iso 639-2
	if len(d.Language) != 3 {
		return errors.InvalidLocalizationLanguageCode
	}
	_, err := language.ParseBase(d.Language)
	if err != nil {
		return errors.InvalidLocalizationLanguageCode
	}
	return nil
}
