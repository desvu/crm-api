package tag

import (
	"context"
	"errors"

	"github.com/qilin/crm-api/internal/domain/entity"
	"github.com/qilin/crm-api/internal/domain/service"
)

type Service struct {
	ServiceParams
}

var ErrTagNotFound = errors.New("tag not found")
var ErrInvalidTagIDs = errors.New("invalid tag ids")

func (s Service) Create(ctx context.Context, data *service.CreateTagData) (*entity.Tag, error) {
	tag := &entity.Tag{
		Name: data.Name,
	}

	if err := s.TagRepository.Create(ctx, tag); err != nil {
		return nil, err
	}

	return tag, nil
}

func (s Service) Update(ctx context.Context, data *service.UpdateTagData) (*entity.Tag, error) {
	tag, err := s.GetExistByID(ctx, data.ID)
	if err != nil {
		return nil, err
	}

	if tag.Name != data.Name {
		tag.Name = data.Name
		if err = s.TagRepository.Update(ctx, tag); err != nil {
			return nil, err
		}
	}

	return tag, nil
}

func (s Service) Delete(ctx context.Context, id uint) error {
	tag, err := s.GetExistByID(ctx, id)
	if err != nil {
		return err
	}

	return s.TagRepository.Delete(ctx, tag)
}

func (s Service) GetByID(ctx context.Context, id uint) (*entity.Tag, error) {
	return s.TagRepository.FindByID(ctx, id)
}

func (s Service) GetExistByID(ctx context.Context, id uint) (*entity.Tag, error) {
	tag, err := s.TagRepository.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if tag == nil {
		return nil, ErrTagNotFound
	}

	return tag, nil
}

func (s Service) GetByIDs(ctx context.Context, ids []uint) ([]entity.Tag, error) {
	return s.TagRepository.FindByIDs(ctx, ids)
}

func (s Service) GetByGameID(ctx context.Context, gameID uint) ([]entity.Tag, error) {
	gameTags, err := s.GameTagRepository.FindByGameID(ctx, gameID)
	if err != nil {
		return nil, err
	}

	return s.GetByIDs(ctx, entity.NewGameTagArray(gameTags).IDs())
}

func (s Service) UpdateTagsForGame(ctx context.Context, game *entity.Game, tagIDs []uint) error {
	tags, err := s.GetByIDs(ctx, tagIDs)
	if err != nil {
		return err
	}

	// checking for IDs among the tags
	if len(tags) != len(tagIDs) {
		return ErrInvalidTagIDs
	}

	currentGameTags, err := s.GameTagRepository.FindByGameID(ctx, game.ID)
	if err != nil {
		return err
	}

	err = s.GameTagRepository.DeleteMultiple(ctx, s.getGameTagsForDelete(tagIDs, currentGameTags))
	if err != nil {
		return err
	}

	err = s.GameTagRepository.CreateMultiple(ctx, s.getGameTagsForInsert(game.ID, tagIDs, currentGameTags))
	if err != nil {
		return err
	}

	return nil
}

func (s Service) getGameTagsForInsert(gameID uint, newTagIDs []uint, currentGameTags []entity.GameTag) []entity.GameTag {
	gameTags := make([]entity.GameTag, 0)
	for _, newTagID := range newTagIDs {
		var hasMatch bool
		for _, currentGameTag := range currentGameTags {
			if newTagID == currentGameTag.TagID {
				hasMatch = true
			}
		}

		if !hasMatch {
			gameTags = append(gameTags, entity.GameTag{
				GameID: gameID,
				TagID:  newTagID,
			})
		}
	}

	return gameTags
}

func (s Service) getGameTagsForDelete(newTagIDs []uint, currentGameTags []entity.GameTag) []entity.GameTag {
	gameTags := make([]entity.GameTag, 0)
	for _, currentGameTag := range currentGameTags {
		var hasMatch bool
		for _, newTagID := range newTagIDs {
			if currentGameTag.TagID == newTagID {
				hasMatch = true
			}
		}

		if !hasMatch {
			gameTags = append(gameTags, entity.GameTag{
				ID:     currentGameTag.ID,
				GameID: currentGameTag.GameID,
				TagID:  currentGameTag.TagID,
			})
		}
	}

	return gameTags
}
