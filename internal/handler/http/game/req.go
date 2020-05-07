package game

import (
	"time"

	"github.com/labstack/echo/v4"
	"github.com/qilin/crm-api/internal/domain/entity"
	gameenum "github.com/qilin/crm-api/internal/domain/enum/game"
	"github.com/qilin/crm-api/internal/domain/service"
)

type reqUpsert struct {
	ID          *string      `json:"id"`
	Title       *string      `json:"title"`
	Type        *string      `json:"type"`
	Slug        *string      `json:"slug"`
	Summary     *string      `json:"summary"`
	Description *string      `json:"description"`
	License     *string      `json:"license"`
	Platforms   *[]string    `json:"platforms"`
	Developers  *[]uint      `json:"developers"`
	Features    *[]uint      `json:"features"`
	Genres      *[]uint      `json:"genres"`
	Publishers  *[]uint      `json:"publishers"`
	Tags        *[]uint      `json:"tags"`
	Media       *[]uint      `json:"media"`
	ReleaseDate *time.Time   `json:"release_date"`
	SocialLinks []socialLink `json:"social_links"`
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
			Tags:        req.Tags,
			Developers:  req.Developers,
			Publishers:  req.Publishers,
			Features:    req.Features,
			Genres:      req.Genres,
			Media:       req.Media,
			Platforms:   nil,
			ReleaseDate: req.ReleaseDate,
			SocialLinks: convertSocialLinksToServiceSocialLinks(req.SocialLinks),
		},
	}, nil
}

func convertSocialLinksToServiceSocialLinks(links []socialLink) *[]service.SocialLink {
	list := make([]service.SocialLink, len(links))
	for i, item := range links {
		list[i] = service.SocialLink{URL: item.URL}
	}
	return &list
}

func convertEntitySocialLinksToSocialLinks(links []entity.SocialLink) []socialLink {
	list := make([]socialLink, len(links))
	for i, item := range links {
		list[i] = socialLink{URL: item.URL}
	}
	return list
}
