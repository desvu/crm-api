package game

import (
	"time"

	"github.com/labstack/echo/v4"
	gameenum "github.com/qilin/crm-api/internal/domain/enum/game"
	"github.com/qilin/crm-api/internal/domain/service"
)

type reqUpsert struct {
	ID          *string    `json:"id"`
	Title       *string    `json:"title"`
	Type        *string    `json:"type"`
	Slug        *string    `json:"slug"`
	Summary     *string    `json:"summary"`
	Description *string    `json:"description"`
	License     *string    `json:"license"`
	Trailer     *string    `json:"trailer"`
	Platforms   *[]string  `json:"platforms"`
	Developers  *[]uint    `json:"developers"`
	Features    *[]uint    `json:"features"`
	Genres      *[]uint    `json:"genres"`
	Publishers  *[]uint    `json:"publishers"`
	Tags        *[]uint    `json:"tags"`
	Media       *[]uint    `json:"media"`
	ReleaseDate *time.Time `json:"release_date"`
}

func convertUpsertRequest(c echo.Context) (*service.UpsertGameData, error) {
	req := new(reqUpsert)
	if err := c.Bind(req); err != nil {
		return nil, err
	}

	return &service.UpsertGameData{
		ID:    req.ID,
		Title: req.Title,
		Slug:  req.Slug,
		Type:  gameenum.NewTypePointerByStringPointer(req.Type),
		CommonGameData: service.CommonGameData{
			Summary:     req.Summary,
			Description: req.Description,
			License:     req.License,
			Trailer:     req.Trailer,
			Tags:        req.Tags,
			Developers:  req.Developers,
			Publishers:  req.Publishers,
			Features:    req.Features,
			Genres:      req.Genres,
			Media:       req.Media,
			Platforms:   nil,
			ReleaseDate: req.ReleaseDate,
		},
	}, nil
}
