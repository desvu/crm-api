package graphql

import (
	"strconv"

	"github.com/qilin/crm-api/internal/domain/entity"
	"github.com/qilin/crm-api/internal/domain/enum/game"
)

type gameConverter struct {
}

var ISO8601Extended = "2006-01-02T15:04:05.000Z"

func (c gameConverter) convertGame(g *entity.GameEx) *Game {
	if g.Revision == nil {
		return nil
	}

	return &Game{
		ID: g.ID,
		Revision: &Revision{
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

func (c gameConverter) convertDeveloperArray(items []entity.Developer) []*Developer {
	var result []*Developer
	for _, developer := range items {
		result = append(result, &Developer{
			ID:   strconv.Itoa(int(developer.ID)),
			Name: developer.Name,
		})
	}

	return result
}

func (c gameConverter) convertPublisherArray(items []entity.Publisher) []*Publisher {
	var result []*Publisher
	for _, publisher := range items {
		result = append(result, &Publisher{
			ID:   strconv.Itoa(int(publisher.ID)),
			Name: publisher.Name,
		})
	}

	return result
}

func (c gameConverter) convertTagArray(items []entity.Tag) []*Tag {
	var result []*Tag
	for _, tag := range items {
		result = append(result, &Tag{
			ID:   strconv.Itoa(int(tag.ID)),
			Name: tag.Name,
		})
	}

	return result
}

func (c gameConverter) convertFeatureArray(items []entity.Feature) []*Feature {
	var result []*Feature
	for _, feature := range items {
		result = append(result, &Feature{
			ID:   strconv.Itoa(int(feature.ID)),
			Name: feature.Name,
		})
	}

	return result
}

func (c gameConverter) convertGenreArray(items []entity.Genre) []*Genre {
	var result []*Genre
	for _, genre := range items {
		result = append(result, &Genre{
			ID:   strconv.Itoa(int(genre.ID)),
			Name: genre.Name,
		})
	}

	return result
}

func (c gameConverter) convertGameType(t game.Type) GameType {
	switch t {
	case game.TypeDesktop:
		return GameTypeDesktop
	case game.TypeWeb:
		return GameTypeWeb
	}
	return GameTypeDesktop
}

func (c gameConverter) convertPlatforms(platforms ...game.Platform) []Platform {
	var res = make([]Platform, len(platforms))
	for i, p := range platforms {
		res[i] = c.convertPlatform(p)
	}
	return res
}

func (c gameConverter) convertPlatform(p game.Platform) Platform {
	switch p {
	case game.PlatformLinux:
		return PlatformLinux
	case game.PlatformWeb:
		return PlatformWeb
	case game.PlatformWindows:
		return PlatformWindows
	case game.PlatformMacOS:
		return PlatformMacOs
	}
	return PlatformWindows
}
