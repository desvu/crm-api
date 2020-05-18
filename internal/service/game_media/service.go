package game_media

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"math"

	"github.com/disintegration/imaging"
	"github.com/google/uuid"
	"github.com/h2non/filetype"
	"github.com/qilin/crm-api/internal/domain/entity"
	"github.com/qilin/crm-api/internal/domain/enum/game_media"
	"github.com/qilin/crm-api/internal/domain/errors"
	"github.com/qilin/crm-api/internal/domain/service"
)

type Service struct {
	ServiceParams
}

func (s *Service) Create(ctx context.Context, data *service.CreateGameMediaData) (*entity.GameMedia, error) {
	gameMedia := &entity.GameMedia{
		GameID:   data.Game.ID,
		Type:     data.Type,
		FilePath: fmt.Sprintf("/game/%s/media/%s.jpg", data.Game.ID, uuid.New().String()),
	}

	if err := s.GameMediaRepository.Create(ctx, gameMedia); err != nil {
		return nil, err
	}

	return gameMedia, nil
}

func (s *Service) Upload(ctx context.Context, data *service.UploadGameMediaData) (*entity.GameMedia, error) {
	gameMedia, err := s.GetByID(ctx, data.ID)
	if err != nil {
		return nil, err
	}

	src, err := getResultImage(data.Image, gameMedia.Type)
	if err != nil {
		return nil, err
	}

	w, err := s.Env.Storage.Bucket.NewWriter(ctx, gameMedia.FilePath, nil)
	if err != nil {
		return nil, err
	}

	_, err = w.Write(src)
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

func (s *Service) Delete(ctx context.Context, id uint) error {
	cover, err := s.GetByID(ctx, id)
	if err != nil {
		return err
	}

	if err := s.Env.Storage.Bucket.Delete(ctx, cover.FilePath); err != nil {
		return err
	}

	return s.GameMediaRepository.Delete(ctx, cover)
}

func (s *Service) GetByID(ctx context.Context, id uint) (*entity.GameMedia, error) {
	cover, err := s.GameMediaRepository.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if cover == nil {
		return nil, errors.MediaNotFound
	}

	return cover, nil
}

func (s *Service) GetByIDs(ctx context.Context, ids []uint) ([]entity.GameMedia, error) {
	return s.GameMediaRepository.FindByIDs(ctx, ids)
}

func (s *Service) GetByRevision(ctx context.Context, revision *entity.GameRevision) ([]entity.GameMedia, error) {
	revisionMedia, err := s.GameRevisionMediaRepository.FindByRevisionID(ctx, revision.ID)
	if err != nil {
		return nil, err
	}

	if len(revisionMedia) == 0 {
		return nil, nil
	}

	return s.GetByIDs(ctx, entity.NewGameRevisionMediaArray(revisionMedia).MediaIDs())
}

func (s *Service) UpdateForGameRevision(ctx context.Context, gameRevision *entity.GameRevision, mediaIDs []uint) error {
	media, err := s.GetByIDs(ctx, mediaIDs)
	if err != nil {
		return err
	}

	// checking for IDs among the media
	if len(media) != len(mediaIDs) {
		return errors.InvalidMediaIDs
	}

	currentGameMedia, err := s.GameRevisionMediaRepository.FindByRevisionID(ctx, gameRevision.ID)
	if err != nil {
		return err
	}

	err = s.GameRevisionMediaRepository.DeleteMultiple(ctx, getGameMediaForDelete(mediaIDs, currentGameMedia))
	if err != nil {
		return err
	}

	err = s.GameRevisionMediaRepository.CreateMultiple(ctx, getGameMediaForInsert(gameRevision.ID, mediaIDs, currentGameMedia))
	if err != nil {
		return err
	}

	return nil
}

func getGameMediaForInsert(revisionID uint, newMediaIDs []uint, currentGameMedia []entity.GameRevisionMedia) []entity.GameRevisionMedia {
	gameMedia := make([]entity.GameRevisionMedia, 0)
	for _, newMediaID := range newMediaIDs {
		var hasMatch bool
		for _, currentGameMedia := range currentGameMedia {
			if newMediaID == currentGameMedia.MediaID {
				hasMatch = true
			}
		}

		if !hasMatch {
			gameMedia = append(gameMedia, entity.GameRevisionMedia{
				RevisionID: revisionID,
				MediaID:    newMediaID,
			})
		}
	}

	return gameMedia
}

func getGameMediaForDelete(newMediaIDs []uint, currentGameMedia []entity.GameRevisionMedia) []entity.GameRevisionMedia {
	gameMedia := make([]entity.GameRevisionMedia, 0)
	for _, currentGameMedia := range currentGameMedia {
		var hasMatch bool
		for _, newMediaID := range newMediaIDs {
			if currentGameMedia.MediaID == newMediaID {
				hasMatch = true
			}
		}

		if !hasMatch {
			gameMedia = append(gameMedia, entity.GameRevisionMedia{
				ID:         currentGameMedia.ID,
				RevisionID: currentGameMedia.RevisionID,
				MediaID:    currentGameMedia.MediaID,
			})
		}
	}

	return gameMedia
}

func getResultImage(src []byte, t game_media.Type) ([]byte, error) {
	kind, err := filetype.Match(src)
	if err != nil {
		return nil, err
	}

	if kind.MIME.Value != "image/png" {
		return nil, errors.InvalidMediaMIMEType
	}

	img, err := imaging.Decode(bytes.NewReader(src))
	if err != nil {
		return nil, err
	}

	if !checkResolution(img.Bounds().Dx(), img.Bounds().Dy(), t) {
		return nil, errors.InvalidMediaResolution
	}

	if !checkAspectRatio(img.Bounds().Dx(), img.Bounds().Dy(), t) {
		return nil, errors.InvalidMediaAspectRatio
	}

	if !t.IsNeedResize {
		return src, nil
	}

	buf := new(bytes.Buffer)
	err = imaging.Encode(buf, imaging.Resize(img, t.ResultWidth, t.ResultHeight, imaging.Lanczos), imaging.JPEG)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func checkResolution(w, h int, t game_media.Type) bool {
	if !t.IsNeedValidate {
		return true
	}

	if w < t.ResultWidth || h != t.ResultHeight {
		return false
	}

	return true
}

func checkAspectRatio(w, h int, t game_media.Type) bool {
	if !t.IsNeedValidate {
		return true
	}

	if w == int(math.Round(float64(h)/float64(t.ResultHeight)*float64(t.ResultWidth))) {
		log.Println(w, int(math.Round(float64(h)/float64(t.ResultHeight)*float64(t.ResultWidth))))
		return true
	}

	return false
}
