package game

import (
	"context"

	"github.com/qilin/crm-api/internal/domain/entity"
	"github.com/qilin/crm-api/internal/domain/service"
	"github.com/qilin/crm-api/pkg/errors/grpcerror"

	"github.com/qilin/crm-api/pkg/grpc/proto"
)

func (h *Handler) GetBySlug(ctx context.Context, request *proto.GetBySlugRequest) (*proto.GameResponse, error) {
	game, err := h.GameService.GetExBySlug(ctx, request.Slug)
	if err != nil {
		return nil, grpcerror.New(err)
	}

	result, err := h.convertGame(game)
	if err != nil {
		return nil, grpcerror.New(err)
	}

	return &proto.GameResponse{Game: result}, nil
}

func (h *Handler) GetByID(ctx context.Context, request *proto.GetByIDRequest) (*proto.GameResponse, error) {
	game, err := h.GameService.GetExLastPublishedByID(ctx, request.GameID)
	if err != nil {
		return nil, grpcerror.New(err)
	}

	result, err := h.convertGame(game)
	if err != nil {
		return nil, grpcerror.New(err)
	}

	return &proto.GameResponse{Game: result}, nil
}

func (h *Handler) GetByIDAndRevisionID(ctx context.Context, request *proto.Request) (*proto.GameResponse, error) {
	game, err := h.GameService.GetExByIDAndRevisionID(ctx, request.GameID, uint(request.RevisionID))
	if err != nil {
		return nil, grpcerror.New(err)
	}

	result, err := h.convertGame(game)
	if err != nil {
		return nil, grpcerror.New(err)
	}

	return &proto.GameResponse{Game: result}, nil
}

func (h *Handler) SearchByTitle(ctx context.Context, request *proto.FindByTitleRequest) (*proto.SearchByTitleResponse, error) {
	games, err := h.GameService.GetByTitleSubstring(ctx, service.GetByTitleSubstringData{
		Title:  request.Title,
		Limit:  uint(request.Limit),
		Offset: uint(request.Offset),
	})
	if err != nil {
		return nil, err
	}

	response := make([]*proto.Game, len(games))
	for i, g := range games {
		response[i], err = h.convertGame(&g)
		if err != nil {
			return nil, err
		}
	}

	return &proto.SearchByTitleResponse{
		Games: response,
	}, nil
}

func (h *Handler) convertGame(game *entity.GameEx) (*proto.Game, error) {
	result := &proto.Game{
		ID:          game.ID,
		Title:       game.Title,
		Slug:        game.Slug,
		Type:        game.Type.String(),
		RevisionID:  uint64(game.Revision.ID),
		Summary:     game.Revision.Summary,
		Description: game.Revision.Description,
		License:     game.Revision.License,
		Platforms:   game.Revision.Platforms.Strings(),
		Trailer:     game.Revision.Trailer,
		PlayTime:    uint32(game.Revision.PlayTime),
	}

	for _, item := range game.Revision.Tags {
		result.Tags = append(result.Tags, &proto.Tag{
			ID:   uint64(item.ID),
			Name: item.Name,
		})
	}

	for _, item := range game.Revision.Developers {
		result.Developers = append(result.Developers, &proto.Developer{
			ID:   uint64(item.ID),
			Name: item.Name,
		})
	}

	for _, item := range game.Revision.Publishers {
		result.Publishers = append(result.Publishers, &proto.Publisher{
			ID:   uint64(item.ID),
			Name: item.Name,
		})
	}

	for _, item := range game.Revision.Genres {
		result.Genres = append(result.Genres, &proto.Genre{
			ID:   uint64(item.ID),
			Name: item.Name,
		})
	}

	for _, item := range game.Revision.Features {
		result.Features = append(result.Features, &proto.Feature{
			ID:   uint64(item.ID),
			Name: item.Name,
			Icon: item.Icon.String(),
		})
	}

	for _, item := range game.Revision.Media {
		result.Media = append(result.Media, &proto.Media{
			Type: item.Type.Name,
			Url:  h.URLBuilder.BuildGameMedia(&item),
		})
	}

	for _, item := range game.Revision.Localization {
		result.Localizations = append(result.Localizations, &proto.Localization{
			Language:  item.Language,
			Interface: item.Interface,
			Audio:     item.Audio,
			Subtitles: item.Subtitles,
		})
	}

	for _, item := range game.Revision.SocialLinks {
		result.SocialLinks = append(result.SocialLinks, &proto.SocialLink{
			Type: item.Type.String(),
			Url:  item.URL,
		})
	}

	for _, item := range game.Revision.SystemRequirements {
		r := &proto.SystemRequirements{
			Platform: item.Platform.String(),
		}
		if item.Minimal != nil {
			r.Minimal = &proto.RequirementsSet{
				CPU:       item.Minimal.CPU,
				GPU:       item.Minimal.GPU,
				DiskSpace: uint32(item.Minimal.DiskSpace),
				RAM:       uint32(item.Minimal.RAM),
			}
		}
		if item.Recommended != nil {
			r.Recommended = &proto.RequirementsSet{
				CPU:       item.Recommended.CPU,
				GPU:       item.Recommended.GPU,
				DiskSpace: uint32(item.Recommended.DiskSpace),
				RAM:       uint32(item.Recommended.RAM),
			}
		}

		result.SystemRequirements = append(result.SystemRequirements, r)
	}

	for _, item := range game.Revision.Rating {
		result.Rating = append(result.Rating, &proto.Rating{
			Agency:              item.Agency.String(),
			Rating:              item.Rating.String(),
			DisplayOnlineNotice: item.DisplayOnlineNotice,
			ShowAgeRestrict:     item.ShowAgeRestrict,
		})
	}

	for _, item := range game.Revision.Review {
		result.Reviews = append(result.Reviews, &proto.Review{
			PressName: item.PressName,
			Link:      item.Link,
			Score:     item.Score,
			Quote:     item.Quote,
		})
	}

	return result, nil
}
