package localization

import (
	"context"

	"github.com/qilin/crm-api/internal/domain/entity"
	"github.com/qilin/crm-api/internal/domain/errors"
	"github.com/qilin/crm-api/internal/domain/service"
)

type Service struct {
	ServiceParams
}

func (s *Service) GetByID(ctx context.Context, id uint) (*entity.Localization, error) {
	return s.LocalizationRepository.FindByID(ctx, id)
}

func (s *Service) GetExistByID(ctx context.Context, id uint) (*entity.Localization, error) {
	genre, err := s.LocalizationRepository.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if genre == nil {
		return nil, errors.LocalizationNotFound
	}

	return genre, nil
}

func (s *Service) GetByIDs(ctx context.Context, ids []uint) ([]entity.Localization, error) {
	return s.LocalizationRepository.FindByIDs(ctx, ids)
}

func (s *Service) GetByGameRevisionID(ctx context.Context, gameID uint) ([]entity.Localization, error) {
	return s.LocalizationRepository.FindByGameRevisionID(ctx, gameID)
}

func (s *Service) UpdateLocalizationsForGameRevision(ctx context.Context, gameRevision *entity.GameRevision, localizations []service.LocalizationData) error {
	langs := make([]string, len(localizations))
	for i, l := range localizations {
		langs[i] = l.Language
	}

	locs, err := s.LocalizationRepository.FindByGameRevisionIDAndLanguage(ctx, gameRevision.ID, langs)
	if err != nil {
		return err
	}

	currentGameLocalizations, err := s.LocalizationRepository.FindByGameRevisionID(ctx, gameRevision.ID)
	if err != nil {
		return err
	}

	err = s.LocalizationRepository.DeleteMultiple(ctx, getGameLocalizationsForDelete(locs, currentGameLocalizations))
	if err != nil {
		return err
	}

	err = s.LocalizationRepository.CreateMultiple(ctx, getGameLocalizationsForInsert(gameRevision.ID, localizations, currentGameLocalizations))
	if err != nil {
		return err
	}

	err = s.LocalizationRepository.UpdateMultiple(ctx, getGameLocalizationsForUpdate(gameRevision.ID, localizations, currentGameLocalizations))
	if err != nil {
		return err
	}

	return nil
}

func getGameLocalizationsForInsert(gameID uint, newLocalizations []service.LocalizationData, currentGameLocalizations []entity.Localization) []entity.Localization {
	gameLocalizations := make([]entity.Localization, 0)
	for _, newLocalization := range newLocalizations {
		var hasMatch bool
		for _, currentGameLocalization := range currentGameLocalizations {
			if newLocalization.Language == currentGameLocalization.Language {
				hasMatch = true
			}
		}

		if !hasMatch {
			gameLocalizations = append(gameLocalizations, entity.Localization{
				GameRevisionID: gameID,
				Language:       newLocalization.Language,
				Interface:      newLocalization.Interface,
				Audio:          newLocalization.Audio,
				Subtitles:      newLocalization.Subtitles,
			})
		}
	}
	return gameLocalizations
}

func getGameLocalizationsForUpdate(gameID uint, newLocalizations []service.LocalizationData, currentGameLocalizations []entity.Localization) []entity.Localization {
	gameLocalizations := make([]entity.Localization, 0)
	for _, newLocalization := range newLocalizations {
		for _, currentGameLocalization := range currentGameLocalizations {
			if newLocalization.Language != currentGameLocalization.Language {
				continue
			}

			if newLocalization.Interface != currentGameLocalization.Interface || newLocalization.Audio != currentGameLocalization.Audio || newLocalization.Subtitles != currentGameLocalization.Subtitles {
				gameLocalizations = append(gameLocalizations, entity.Localization{
					ID:             currentGameLocalization.ID,
					GameRevisionID: gameID,
					Language:       newLocalization.Language,
					Interface:      newLocalization.Interface,
					Audio:          newLocalization.Audio,
					Subtitles:      newLocalization.Subtitles,
				})
			}
		}
	}
	return gameLocalizations
}

func getGameLocalizationsForDelete(localizations []entity.Localization, currentGameLocalizations []entity.Localization) []entity.Localization {

	gameLocalizations := make([]entity.Localization, 0)
	for _, currentGameLocalization := range currentGameLocalizations {
		var hasMatch bool
		for _, newLocalization := range localizations {
			if currentGameLocalization.Language == newLocalization.Language {
				hasMatch = true
			}
		}

		if !hasMatch {
			gameLocalizations = append(gameLocalizations, entity.Localization{
				ID:             currentGameLocalization.ID,
				GameRevisionID: currentGameLocalization.GameRevisionID,
				Language:       currentGameLocalization.Language,
				Interface:      currentGameLocalization.Interface,
				Audio:          currentGameLocalization.Audio,
				Subtitles:      currentGameLocalization.Subtitles,
			})
		}
	}
	return gameLocalizations
}
