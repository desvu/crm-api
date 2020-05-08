package game

import (
	"time"

	"github.com/labstack/echo/v4"
	gameenum "github.com/qilin/crm-api/internal/domain/enum/game"
	"github.com/qilin/crm-api/internal/domain/service"
)

type reqUpsert struct {
	ID           *string        `json:"id"`
	Title        *string        `json:"title"`
	Type         *string        `json:"type"`
	Slug         *string        `json:"slug"`
	Summary      *string        `json:"summary"`
	Description  *string        `json:"description"`
	License      *string        `json:"license"`
	Platforms    *[]string      `json:"platforms"`
	Developers   *[]uint        `json:"developers"`
	Features     *[]uint        `json:"features"`
	Genres       *[]uint        `json:"genres"`
	Publishers   *[]uint        `json:"publishers"`
	Tags         *[]uint        `json:"tags"`
	Media        *[]uint        `json:"media"`
	ReleaseDate  *time.Time     `json:"release_date"`
	Localization []localization `json:"localization"`
}

func convertUpsertRequest(c echo.Context) (*service.UpsertGameData, error) {
	req := new(reqUpsert)
	if err := c.Bind(req); err != nil {
		return nil, err
	}

	data := &service.UpsertGameData{
		ID:    req.ID,
		Title: req.Title,
		Slug:  req.Slug,
		Type:  gameenum.NewTypePointerByStringPointer(req.Type),
		CommonGameData: service.CommonGameData{
			Summary:     req.Summary,
			Description: req.Description,
			License:     req.License,
			Tags:        req.Tags,
			Developers:  req.Developers,
			Publishers:  req.Publishers,
			Features:    req.Features,
			Genres:      req.Genres,
			Media:       req.Media,
			Platforms:   nil,
			ReleaseDate: req.ReleaseDate,
		},
	}

	if len(req.Localization) > 0 {
		localizations := make([]service.LocalizationData, len(req.Localization))
		for i, l := range req.Localization {
			localizations[i] = service.LocalizationData{
				Language:  l.Language,
				Interface: l.Interface,
				Audio:     l.Audio,
				Subtitles: l.Subtitles,
			}
		}
		data.CommonGameData.Localizations = &localizations
	}

	return data, nil
}
