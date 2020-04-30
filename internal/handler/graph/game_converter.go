package graph

import (
	"strconv"

	"github.com/qilin/crm-api/internal/handler/graph/model"

	"github.com/qilin/crm-api/internal/domain/entity"
	"github.com/qilin/crm-api/internal/domain/enum/game"
)

type gameConverter struct {
}

var ISO8601Extended = "2006-01-02T15:04:05.000Z"

func (c gameConverter) convertGame(g *entity.GameEx) *model.Game {
	if g.Revision == nil {
		return nil
	}

	return &model.Game{
		ID: g.ID,
		Revision: &model.Revision{
			ID:          strconv.Itoa(int(g.Revision.ID)),
			GameID:      g.Revision.GameID,
			Summary:     g.Revision.Summary,
			Description: g.Revision.Description,
			License:     g.Revision.License,
			Developers:  c.convertDeveloperArray(g.Revision.Developers),
			Publishers:  c.convertPublisherArray(g.Revision.Publishers),
			Genres:      c.convertGenreArray(g.Revision.Genres),
			Tags:        c.convertTagArray(g.Revision.Tags),
			Features:    c.convertFeatureArray(g.Revision.Features),
			Platforms:   c.convertPlatforms(g.Revision.Platforms...),
			ReleaseDate: g.Revision.ReleaseDate.Format(ISO8601Extended),
		},
	}
}

func (c gameConverter) convertDeveloperArray(items []entity.Developer) []*model.Developer {
	var result []*model.Developer
	for _, developer := range items {
		result = append(result, &model.Developer{
			ID:   strconv.Itoa(int(developer.ID)),
			Name: developer.Name,
		})
	}

	return result
}

func (c gameConverter) convertPublisherArray(items []entity.Publisher) []*model.Publisher {
	var result []*model.Publisher
	for _, publisher := range items {
		result = append(result, &model.Publisher{
			ID:   strconv.Itoa(int(publisher.ID)),
			Name: publisher.Name,
		})
	}

	return result
}

func (c gameConverter) convertTagArray(items []entity.Tag) []*model.Tag {
	var result []*model.Tag
	for _, tag := range items {
		result = append(result, &model.Tag{
			ID:   strconv.Itoa(int(tag.ID)),
			Name: tag.Name,
		})
	}

	return result
}

func (c gameConverter) convertFeatureArray(items []entity.Feature) []*model.Feature {
	var result []*model.Feature
	for _, feature := range items {
		result = append(result, c.convertFeature(feature))
	}

	return result
}

func (c gameConverter) convertFeature(feature entity.Feature) *model.Feature {
	return &model.Feature{
		ID:   strconv.Itoa(int(feature.ID)),
		Name: feature.Name,
		Icon: feature.Icon.String(),
	}
}

func (c gameConverter) convertGenreArray(items []entity.Genre) []*model.Genre {
	var result []*model.Genre
	for _, genre := range items {
		result = append(result, &model.Genre{
			ID:   strconv.Itoa(int(genre.ID)),
			Name: genre.Name,
		})
	}

	return result
}

func (c gameConverter) convertGameType(t game.Type) model.GameType {
	switch t {
	case game.TypeDesktop:
		return model.GameTypeDesktop
	case game.TypeWeb:
		return model.GameTypeWeb
	}
	return model.GameTypeDesktop
}

func (c gameConverter) convertPlatforms(platforms ...game.Platform) []model.GamePlatform {
	var res = make([]model.GamePlatform, len(platforms))
	for i, p := range platforms {
		res[i] = c.convertPlatform(p)
	}
	return res
}

func (c gameConverter) convertPlatform(p game.Platform) model.GamePlatform {
	switch p {
	case game.PlatformLinux:
		return model.GamePlatformLinux
	case game.PlatformWeb:
		return model.GamePlatformWeb
	case game.PlatformWindows:
		return model.GamePlatformWindows
	case game.PlatformMacOS:
		return model.GamePlatformMacOs
	}
	return model.GamePlatformWindows
}
