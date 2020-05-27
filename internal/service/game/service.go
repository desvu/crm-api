package game

import (
	"context"

	"github.com/google/uuid"
	"github.com/qilin/crm-api/internal/domain/entity"
	"github.com/qilin/crm-api/internal/domain/enum/game_revision"
	"github.com/qilin/crm-api/internal/domain/errors"
	"github.com/qilin/crm-api/internal/domain/repository"
	"github.com/qilin/crm-api/internal/domain/service"
)

type Service struct {
	ServiceParams
}

func (s *Service) Create(ctx context.Context, data *service.CreateGameData) (*entity.GameEx, error) {
	if err := data.Validate(); err != nil {
		return nil, errors.NewValidation(err)
	}

	if err := s.checkNoExistGameBySlug(ctx, data.Slug); err != nil {
		return nil, err
	}

	game := &entity.Game{
		ID:    uuid.New().String(),
		Title: data.Title,
		Type:  data.Type,
		Slug:  data.Slug,
	}

	var updatedRevision *entity.GameRevisionEx
	if err := s.Transactor.Transact(ctx, func(tx context.Context) error {
		if err := s.GameRepository.Create(tx, game); err != nil {
			return err
		}

		revision, err := s.GameRevisionService.GetDraftByGame(tx, game)
		if err != nil {
			return err
		}

		updatedRevision, err = s.GameRevisionService.Update(tx, &service.UpdateGameRevisionData{
			ID:                 revision.ID,
			Summary:            data.Summary,
			Description:        data.Description,
			License:            data.License,
			Trailer:            data.Trailer,
			Tags:               data.Tags,
			Developers:         data.Developers,
			Publishers:         data.Publishers,
			Features:           data.Features,
			Genres:             data.Genres,
			Media:              data.Media,
			ReleaseDate:        data.ReleaseDate,
			Platforms:          data.Platforms,
			SocialLinks:        data.SocialLinks,
			SystemRequirements: data.SystemRequirements,
			Localizations:      data.Localizations,
			Rating:             data.Ratings,
			Reviews:            data.Reviews,
			PlayTime:           data.PlayTime,
		})

		if err != nil {
			return err
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return &entity.GameEx{
		Game:     *game,
		Revision: updatedRevision,
	}, nil
}

func (s *Service) Update(ctx context.Context, data *service.UpdateGameData) (*entity.GameEx, error) {
	if err := data.Validate(); err != nil {
		return nil, errors.NewValidation(err)
	}

	game, err := s.GetByID(ctx, data.ID)
	if err != nil {
		return nil, err
	}

	revision, err := s.GameRevisionService.GetDraftByGame(ctx, game)
	if err != nil {
		return nil, err
	}

	var updatedRevision *entity.GameRevisionEx
	if err := s.Transactor.Transact(ctx, func(ctx context.Context) error {
		if data.Title != nil {
			game.Title = *data.Title
		}

		if data.Slug != nil && *data.Slug != game.Slug {
			if err := s.checkNoExistGameBySlug(ctx, *data.Slug); err != nil {
				return err
			}

			game.Slug = *data.Slug
		}

		if err = s.GameRepository.Update(ctx, game); err != nil {
			return err
		}

		updatedRevision, err = s.GameRevisionService.Update(ctx, &service.UpdateGameRevisionData{
			ID:                 revision.ID,
			Tags:               data.Tags,
			Developers:         data.Developers,
			Publishers:         data.Publishers,
			Features:           data.Features,
			Genres:             data.Genres,
			Media:              data.Media,
			SocialLinks:        data.SocialLinks,
			SystemRequirements: data.SystemRequirements,
			Localizations:      data.Localizations,
			Rating:             data.Ratings,
			Reviews:            data.Reviews,
			PlayTime:           data.PlayTime,
			Platforms:          data.Platforms,
			ReleaseDate:        data.ReleaseDate,
			Trailer:            data.Trailer,
			License:            data.License,
			Description:        data.Description,
			Summary:            data.Summary,
		})

		if err != nil {
			return err
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return &entity.GameEx{
		Game:     *game,
		Revision: updatedRevision,
	}, nil
}

func (s *Service) Upsert(ctx context.Context, data *service.UpsertGameData) (*entity.GameEx, error) {
	if data.ID != nil {
		return s.Update(ctx, &service.UpdateGameData{
			ID:             *data.ID,
			Title:          data.Title,
			Slug:           data.Slug,
			Type:           data.Type,
			CommonGameData: data.CommonGameData,
		})
	}

	d := &service.CreateGameData{
		CommonGameData: data.CommonGameData,
	}

	if data.Title != nil {
		d.Title = *data.Title
	}

	if data.Slug != nil {
		d.Slug = *data.Slug
	}

	if data.Type != nil {
		d.Type = *data.Type
	}

	if data.Trailer != nil {
		d.Trailer = data.Trailer
	}

	if data.PlayTime != nil {
		d.PlayTime = data.PlayTime
	}

	return s.Create(ctx, d)
}

func (s *Service) Delete(ctx context.Context, id string) error {
	game, err := s.GetByID(ctx, id)
	if err != nil {
		return err
	}

	return s.GameRepository.Delete(ctx, game)
}

func (s *Service) Publish(ctx context.Context, id string) error {
	game, err := s.GetExByID(ctx, id)
	if err != nil {
		return err
	}

	//if err = s.GameStorePublisher.Publish(publisher.PublishGameStoreData{Game: game}); err != nil {
	//	return err
	//}

	revisionStatus := game_revision.StatusPublished // TODO publishing -> published
	_, err = s.GameRevisionService.Update(ctx, &service.UpdateGameRevisionData{
		ID:     game.Revision.ID,
		Status: &revisionStatus,
	})
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) GetByID(ctx context.Context, id string) (*entity.Game, error) {
	_, err := uuid.Parse(id)
	if err != nil {
		return nil, errors.InvalidGameID
	}

	game, err := s.GameRepository.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if game == nil {
		return nil, errors.GameNotFound
	}

	return game, nil
}

func (s *Service) GetBySlug(ctx context.Context, slug string) (*entity.Game, error) {
	game, err := s.GameRepository.FindBySlug(ctx, slug)
	if err != nil {
		return nil, err
	}

	if game == nil {
		return nil, errors.GameNotFound
	}

	return game, nil
}

func (s *Service) GetExLastPublishedByID(ctx context.Context, id string) (*entity.GameEx, error) {
	game, err := s.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	revision, err := s.GameRevisionService.GetLastPublishedByGame(ctx, game)
	if err != nil {
		return nil, err
	}

	return &entity.GameEx{
		Game:     *game,
		Revision: revision,
	}, nil
}

func (s *Service) GetExByID(ctx context.Context, id string) (*entity.GameEx, error) {
	game, err := s.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	revision, err := s.GameRevisionService.GetDraftByGame(ctx, game)
	if err != nil {
		return nil, err
	}

	return &entity.GameEx{
		Game:     *game,
		Revision: revision,
	}, nil
}

func (s *Service) GetExBySlug(ctx context.Context, slug string) (*entity.GameEx, error) {
	game, err := s.GetBySlug(ctx, slug)
	if err != nil {
		return nil, err
	}

	revision, err := s.GameRevisionService.GetLastPublishedByGame(ctx, game)
	if err != nil {
		return nil, err
	}

	return &entity.GameEx{
		Game:     *game,
		Revision: revision,
	}, nil
}

func (s *Service) GetExByIDAndRevisionID(ctx context.Context, id string, revisionID uint) (*entity.GameEx, error) {
	game, err := s.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	revision, err := s.GameRevisionService.GetByIDAndGameID(ctx, revisionID, game.ID)
	if err != nil {
		return nil, err
	}

	return &entity.GameEx{
		Game:     *game,
		Revision: revision,
	}, nil
}

func (s *Service) GetExByFilter(ctx context.Context, data *service.GetByFilterGameData) ([]entity.GameEx, error) {
	if data.Limit == 0 {
		data.Limit = 30
	}

	revisions, err := s.GameRevisionService.GetByFilter(ctx, data)
	if err != nil {
		return nil, err
	}

	games, err := s.GameRepository.FindByIDs(ctx, entity.NewGameRevisionExArray(revisions).GameIDs())
	if err != nil {
		return nil, err
	}

	gamesEx := make([]entity.GameEx, len(revisions))
	for i := range revisions {
		for j := range games {
			if games[j].ID == revisions[i].GameID {
				gamesEx[i] = entity.GameEx{
					Revision: &revisions[i],
					Game:     games[j],
				}
			}
		}
	}

	return gamesEx, nil
}

func (s *Service) GetCountByFilter(ctx context.Context, data *service.GetByFilterGameData) (int, error) {
	return s.GameRevisionService.GetCountByFilter(ctx, data)
}

func (s *Service) GetByTitleSubstring(ctx context.Context, data service.GetByTitleSubstringData) ([]entity.GameEx, error) {
	games, err := s.GameRepository.FindByTitleSubstring(ctx, &repository.FindByTitleSubstringData{
		Title:  data.Title,
		Limit:  data.Limit,
		Offset: data.Offset,
	})
	if err != nil {
		return nil, err
	}

	revisions, err := s.GameRevisionService.GetLastByGameIDs(ctx, entity.NewGameArray(games).IDs())
	if err != nil {
		return nil, err
	}

	var gamesEx []entity.GameEx
	for i := range games {
		for j := range revisions {
			if games[i].ID == revisions[j].GameID {
				gamesEx = append(gamesEx, entity.GameEx{
					Game:     games[i],
					Revision: &revisions[j],
				})
			}
		}
	}

	return gamesEx, nil
}

func (s *Service) checkNoExistGameBySlug(ctx context.Context, slug string) error {
	gameSlug, err := s.GetBySlug(ctx, slug)
	if err != nil && err != errors.GameNotFound {
		return err
	}

	if gameSlug != nil {
		return errors.GameSlugAlreadyExist
	}

	return nil
}
