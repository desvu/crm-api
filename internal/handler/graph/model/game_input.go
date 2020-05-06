package model

import (
	"github.com/qilin/crm-api/internal/domain/enum/game"

	"github.com/qilin/crm-api/internal/domain/service"
)

func (i CreateGameInput) Convert() (*service.CreateGameData, error) {
	platformArray := convertGamePlatformsToGamePlatformArrayPointer(i.Platforms)
	gameType := game.NewTypeByString(i.Type.String())

	data := &service.CreateGameData{
		Title:       i.Title,
		Summary:     i.Summary,
		Description: i.Description,
		Slug:        i.Slug,
		License:     i.License,
		Type:        gameType,
		Platforms:   platformArray,
		ReleaseDate: i.ReleaseDate,
	}

	if len(i.Tags) > 0 {
		var tagIDs []uint
		for _, tagID := range i.Tags {
			tagIDs = append(tagIDs, uint(tagID))
		}
		data.Tags = &tagIDs
	}

	if len(i.Developers) > 0 {
		var developerIDs []uint
		for _, developerID := range i.Developers {
			developerIDs = append(developerIDs, uint(developerID))
		}
		data.Developers = &developerIDs
	}

	if len(i.Publishers) > 0 {
		var publisherIDs []uint
		for _, publisherID := range i.Publishers {
			publisherIDs = append(publisherIDs, uint(publisherID))
		}
		data.Publishers = &publisherIDs
	}

	if len(i.Features) > 0 {
		var featureIDs []uint
		for _, featureID := range i.Features {
			featureIDs = append(featureIDs, uint(featureID))
		}
		data.Features = &featureIDs
	}

	if len(i.Genres) > 0 {
		var genreIDs []uint
		for _, genreID := range i.Genres {
			genreIDs = append(genreIDs, uint(genreID))
		}
		data.Genres = &genreIDs
	}

	if len(i.Localizations) > 0 {
		data.Localizations = convertLocalizationInputToLocalizationData(i.Localizations)
	}

	return data, nil
}

func (i UpdateGameInput) Convert() (*service.UpdateGameData, error) {
	data := &service.UpdateGameData{
		ID:    i.ID,
		Title: i.Title,
	}

	if len(i.Tags) > 0 {
		var tagIDs []uint
		for _, tagID := range i.Tags {
			tagIDs = append(tagIDs, uint(tagID))
		}
		data.Tags = &tagIDs
	}

	if len(i.Developers) > 0 {
		var developerIDs []uint
		for _, developerID := range i.Developers {
			developerIDs = append(developerIDs, uint(developerID))
		}
		data.Developers = &developerIDs
	}

	if len(i.Publishers) > 0 {
		var publisherIDs []uint
		for _, publisherID := range i.Publishers {
			publisherIDs = append(publisherIDs, uint(publisherID))
		}
		data.Publishers = &publisherIDs
	}

	if len(i.Features) > 0 {
		var featureIDs []uint
		for _, featureID := range i.Features {
			featureIDs = append(featureIDs, uint(featureID))
		}
		data.Features = &featureIDs
	}

	if len(i.Genres) > 0 {
		var genreIDs []uint
		for _, genreID := range i.Genres {
			genreIDs = append(genreIDs, uint(genreID))
		}
		data.Genres = &genreIDs
	}

	if len(i.Localizations) > 0 {
		data.Localizations = convertLocalizationInputToLocalizationData(i.Localizations)
	}

	return data, nil
}

func convertGamePlatformsToGamePlatformArrayPointer(items []GamePlatform) *game.PlatformArray {
	if len(items) == 0 {
		return nil
	}

	result := game.PlatformArray{}
	for _, platform := range items {
		result.Add(game.NewPlatformByString(platform.String()))
	}

	return &result
}

func convertLocalizationInputToLocalizationData(items []*LocalizationInput) *[]service.LocalizationData {
	var localizations []service.LocalizationData
	for _, localization := range items {
		if localization == nil {
			continue
		}
		localizations = append(localizations, service.LocalizationData{
			Language:  localization.Language,
			Interface: localization.Interface,
			Audio:     localization.Audio,
			Subtitles: localization.Subtitles,
		})
	}
	return &localizations
}
