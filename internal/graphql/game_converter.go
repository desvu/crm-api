package graphql

import (
	"strconv"

	"github.com/qilin/crm-api/internal/domain/entity"
	"github.com/qilin/crm-api/internal/domain/enum/game"
)

type gameConverter struct {
}

var ISO8601Extended = "2006-01-02T15:04:05.000Z"

func (c gameConverter) convertGame(g *entity.Game) *Game {
	return &Game{
		ID:          strconv.FormatUint(uint64(g.ID), 10),
		Title:       g.Title,
		Summary:     g.Summary,
		Description: g.Description,
		Type:        c.convertGameType(g.Type),
		License:     g.License,
		Ranking:     g.Ranking,
		Platforms:   c.convertPlatforms(g.Platforms...),
		ReleaseDate: g.ReleaseDate.Format(ISO8601Extended),
	}
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
