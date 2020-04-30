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
		for _, req := range i.SystemRequirements {
			r := service.SystemRequirements{
				Platform: game.NewPlatformByString(req.Platform.String()),
			}
			if req.Minimal != nil {
				r.Minimal = &service.RequirementsSet{
					CPU:       req.Minimal.CPU,
					GPU:       req.Minimal.Gpu,
					DiskSpace: uint(req.Minimal.DiskSpace),
					RAM:       uint(req.Minimal.RAM),
				}
			}
			if req.Recommended != nil {
				r.Recommended = &service.RequirementsSet{
					CPU:       req.Recommended.CPU,
					GPU:       req.Recommended.Gpu,
					DiskSpace: uint(req.Recommended.DiskSpace),
					RAM:       uint(req.Recommended.RAM),
				}
			}
			systemRequirements = append(systemRequirements, r)
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
		for _, req := range i.SystemRequirements {
			r := service.SystemRequirements{
				Platform: game.NewPlatformByString(req.Platform.String()),
			}
			if req.Minimal != nil {
				r.Minimal = &service.RequirementsSet{
					CPU:       r.Minimal.CPU,
					GPU:       r.Minimal.GPU,
					DiskSpace: r.Minimal.DiskSpace,
					RAM:       r.Minimal.RAM,
				}
			}
			if req.Recommended != nil {
				r.Recommended = &service.RequirementsSet{
					CPU:       r.Recommended.CPU,
					GPU:       r.Recommended.GPU,
					DiskSpace: r.Recommended.DiskSpace,
					RAM:       r.Recommended.RAM,
				}
			}
			systemRequirements = append(systemRequirements, r)
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
