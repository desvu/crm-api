package tag

import (
	"context"

	"github.com/qilin/crm-api/internal/domain/errors"

	"github.com/qilin/crm-api/internal/domain/entity"
	"github.com/qilin/crm-api/internal/domain/service"
)

type Service struct {
	ServiceParams
}

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
		return nil, errors.TagNotFound
	}

	return tag, nil
}

func (s Service) GetByIDs(ctx context.Context, ids []uint) ([]entity.Tag, error) {
	return s.TagRepository.FindByIDs(ctx, ids)
}

func (s Service) GetByGameRevisionID(ctx context.Context, gameID uint) ([]entity.Tag, error) {
	gameTags, err := s.GameRevisionTagRepository.FindByGameRevisionID(ctx, gameID)
	if err != nil {
		return nil, err
	}

	return s.GetByIDs(ctx, entity.NewGameRevisionTagArray(gameTags).IDs())
}

func (s Service) UpdateTagsForGameRevision(ctx context.Context, gameRevision *entity.GameRevision, tagIDs []uint) error {
	tags, err := s.GetByIDs(ctx, tagIDs)
	if err != nil {
		return err
	}

	// checking for IDs among the tags
	if len(tags) != len(tagIDs) {
		return errors.InvalidTagIDs
	}

	currentGameTags, err := s.GameRevisionTagRepository.FindByGameRevisionID(ctx, gameRevision.ID)
	if err != nil {
		return err
	}

	err = s.GameRevisionTagRepository.DeleteMultiple(ctx, getGameTagsForDelete(tagIDs, currentGameTags))
	if err != nil {
		return err
	}

	err = s.GameRevisionTagRepository.CreateMultiple(ctx, getGameTagsForInsert(gameRevision.ID, tagIDs, currentGameTags))
	if err != nil {
		return err
	}

	return nil
}

func getGameTagsForInsert(gameID uint, newTagIDs []uint, currentGameTags []entity.GameRevisionTag) []entity.GameRevisionTag {
	gameTags := make([]entity.GameRevisionTag, 0)
	for _, newTagID := range newTagIDs {
		var hasMatch bool
		for _, currentGameTag := range currentGameTags {
			if newTagID == currentGameTag.TagID {
				hasMatch = true
			}
		}

		if !hasMatch {
			gameTags = append(gameTags, entity.GameRevisionTag{
				GameRevisionID: gameID,
				TagID:          newTagID,
			})
		}
	}

	return gameTags
}

func getGameTagsForDelete(newTagIDs []uint, currentGameTags []entity.GameRevisionTag) []entity.GameRevisionTag {
	gameTags := make([]entity.GameRevisionTag, 0)
	for _, currentGameTag := range currentGameTags {
		var hasMatch bool
		for _, newTagID := range newTagIDs {
			if currentGameTag.TagID == newTagID {
				hasMatch = true
			}
		}

		if !hasMatch {
			gameTags = append(gameTags, entity.GameRevisionTag{
				ID:             currentGameTag.ID,
				GameRevisionID: currentGameTag.GameRevisionID,
				TagID:          currentGameTag.TagID,
			})
		}
	}

	return gameTags
}
