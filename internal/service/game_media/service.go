package game_media

import (
	"context"
	"strings"

	"github.com/google/uuid"
	"github.com/qilin/crm-api/internal/domain/entity"
	"github.com/qilin/crm-api/internal/domain/service"
)

type Service struct {
	ServiceParams
}

func (s Service) Create(ctx context.Context, data *service.CreateGameMediaData) (*entity.GameMedia, error) {
	game, err := s.GameService.GetByID(ctx, data.GameID)
	if err != nil {
		return nil, err
	}

	fileName := strings.Join([]string{uuid.New().String(), data.Extension}, ".")
	filePath := strings.Join([]string{"/", "game", game.ID, "media", fileName}, "/")

	gameMedia := &entity.GameMedia{
		GameID:   game.ID,
		Type:     data.Type,
		FilePath: filePath,
	}

	if err = s.GameMediaRepository.Create(ctx, gameMedia); err != nil {
		return nil, err
	}

	return gameMedia, nil
}

func (s Service) Upload(ctx context.Context, data *service.UploadGameMediaData) (*entity.GameMedia, error) {
	gameMedia, err := s.GetByID(ctx, data.ID)
	if err != nil {
		return nil, err
	}

	w, err := s.Env.Storage.Bucket.NewWriter(ctx, gameMedia.FilePath, nil)
	if err != nil {
		return nil, err
	}

	_, err = w.Write(data.Image)
	if err != nil {
		return nil, err
	}

	if err = w.Close(); err != nil {
		return nil, err
	}

	gameMedia.IsUploaded = true
	if err = s.GameMediaRepository.Update(ctx, gameMedia); err != nil {
		return nil, err
	}

	return gameMedia, nil
}

func (s Service) Delete(ctx context.Context, id uint) error {
	cover, err := s.GetByID(ctx, id)
	if err != nil {
		return err
	}

	if err := s.Env.Storage.Bucket.Delete(ctx, cover.FilePath); err != nil {
		return err
	}

	return s.GameMediaRepository.Delete(ctx, cover)
}

func (s Service) GetByID(ctx context.Context, id uint) (*entity.GameMedia, error) {
	cover, err := s.GameMediaRepository.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if cover == nil {
		return nil, service.ErrGameMediaNotFound
	}

	return cover, nil
}

func (s Service) GetByIDs(ctx context.Context, ids []uint) ([]entity.GameMedia, error) {
	return s.GameMediaRepository.FindByIDs(ctx, ids)
}

func (s Service) GetByRevision(ctx context.Context, revision *entity.GameRevision) ([]entity.GameMedia, error) {
	revisionMedia, err := s.GameRevisionMediaService.GetByRevision(ctx, revision)
	if err != nil {
		return nil, err
	}

	if len(revisionMedia) == 0 {
		return nil, nil
	}

	return s.GetByIDs(ctx, entity.NewGameRevisionMediaArray(revisionMedia).MediaIDs())
}
