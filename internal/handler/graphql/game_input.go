package graphql

import (
	"context"

	"github.com/qilin/crm-api/internal/domain/service"
)

func (i CreateGameInput) convert(ctx context.Context) (context.Context, *service.CreateGameData) {
	data := &service.CreateGameData{
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

	return ctx, data
}

func (i UpdateGameInput) convert(ctx context.Context) (context.Context, *service.UpdateGameData) {
	data := &service.UpdateGameData{
		ID:    i.ID,
		Title: &i.Title,
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

	return ctx, data
}
