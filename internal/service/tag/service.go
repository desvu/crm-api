package tag

import (
	"context"

	"github.com/pkg/errors"
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
		return nil, errors.WithStack(ErrTagNotFound)
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

func (s Service) AttachTagsToGame(ctx context.Context, gameID uint, tagIDs []uint) error {
	tagIDs, err := s.sanitizeAttachTags(ctx, gameID, tagIDs)
	if err != nil {
		return err
	}

	newGameTags := make([]entity.GameTag, len(tagIDs))
	for i := range tagIDs {
		newGameTags[i] = entity.GameTag{
			GameID: gameID,
			TagID:  tagIDs[i],
		}
	}

	return s.GameTagRepository.CreateMultiple(ctx, newGameTags)
}

func (s Service) sanitizeAttachTags(ctx context.Context, gameID uint, tagIDs []uint) ([]uint, error) {
	if len(tagIDs) == 0 {
		return nil, nil
	}

	game, err := s.GameService.GetExistByID(ctx, gameID)
	if err != nil {
		return nil, err
	}

	tags, err := s.GetByIDs(ctx, tagIDs)
	if err != nil {
		return nil, err
	}

	// checking for IDs among the tags
	if len(tags) != len(tagIDs) {
		return nil, errors.WithStack(ErrInvalidTagIDs)
	}

	currentTags, err := s.GetByGameID(ctx, game.ID)
	if err != nil {
		return nil, err
	}

	// deleting tag IDs that are already attached to the game
	for i := 0; i < len(tagIDs); i++ {
		var hasMatch bool
		for j := range currentTags {
			if tagIDs[i] == currentTags[j].ID {
				hasMatch = true
			}
		}

		if hasMatch {
			tagIDs = append(tagIDs[:i], tagIDs[i+1:]...)
			i--
		}
	}

	return tagIDs, nil
}

func (s Service) DetachTagsFromGame(ctx context.Context, gameID uint, tagIDs []uint) error {
	tagIDs, err := s.sanitizeDetachTags(ctx, gameID, tagIDs)
	if err != nil {
		return err
	}

	deletedGameTags, err := s.GameTagRepository.FindByGameIDAndTagIDs(ctx, gameID, tagIDs)
	if err != nil {
		return err
	}

	return s.GameTagRepository.DeleteMultiple(ctx, deletedGameTags)
}

func (s Service) sanitizeDetachTags(ctx context.Context, gameID uint, tagIDs []uint) ([]uint, error) {
	if len(tagIDs) == 0 {
		return nil, nil
	}

	game, err := s.GameService.GetExistByID(ctx, gameID)
	if err != nil {
		return nil, err
	}

	tags, err := s.GetByIDs(ctx, tagIDs)
	if err != nil {
		return nil, err
	}

	// checking for IDs among the tags
	if len(tags) != len(tagIDs) {
		return nil, errors.WithStack(ErrInvalidTagIDs)
	}

	currentTags, err := s.GetByGameID(ctx, game.ID)
	if err != nil {
		return nil, err
	}

	// deleting tag IDs that are not attached to the game
	for i := 0; i < len(tagIDs); i++ {
		var hasMatch bool
		for j := range currentTags {
			if tagIDs[i] == currentTags[j].ID {
				hasMatch = true
			}
		}

		if !hasMatch {
			tagIDs = append(tagIDs[:i], tagIDs[i+1:]...)
			i--
		}
	}

	return tagIDs, nil
}
