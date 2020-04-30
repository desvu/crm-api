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

	if len(i.SystemRequirements) > 0 {
		var systemRequirements []service.SystemRequirements
		for _, item := range i.SystemRequirements {
			requirementsSet := service.SystemRequirements{
				Platform: game.NewPlatformByString(item.Platform.String()),
			}
			if item.Minimal != nil {
				requirementsSet.Minimal = &service.RequirementsSet{
					CPU:       item.Minimal.CPU,
					GPU:       item.Minimal.Gpu,
					DiskSpace: uint(item.Minimal.DiskSpace),
					RAM:       uint(item.Minimal.RAM),
				}
			}
			if item.Recommended != nil {
				requirementsSet.Recommended = &service.RequirementsSet{
					CPU:       item.Recommended.CPU,
					GPU:       item.Recommended.Gpu,
					DiskSpace: uint(item.Recommended.DiskSpace),
					RAM:       uint(item.Recommended.RAM),
				}
			}
			systemRequirements = append(systemRequirements, requirementsSet)
		}
		data.SystemRequirements = &systemRequirements
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

	if len(i.SystemRequirements) > 0 {
		var systemRequirements []service.SystemRequirements
		for _, item := range i.SystemRequirements {
			requirementsSet := service.SystemRequirements{
				Platform: game.NewPlatformByString(item.Platform.String()),
			}
			if item.Minimal != nil {
				requirementsSet.Minimal = &service.RequirementsSet{
					CPU:       requirementsSet.Minimal.CPU,
					GPU:       requirementsSet.Minimal.GPU,
					DiskSpace: requirementsSet.Minimal.DiskSpace,
					RAM:       requirementsSet.Minimal.RAM,
				}
			}
			if item.Recommended != nil {
				requirementsSet.Recommended = &service.RequirementsSet{
					CPU:       requirementsSet.Recommended.CPU,
					GPU:       requirementsSet.Recommended.GPU,
					DiskSpace: requirementsSet.Recommended.DiskSpace,
					RAM:       requirementsSet.Recommended.RAM,
				}
			}
			systemRequirements = append(systemRequirements, requirementsSet)
		}
		data.SystemRequirements = &systemRequirements
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
