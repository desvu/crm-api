package game

import (
	"time"

	"github.com/labstack/echo/v4"
	"github.com/qilin/crm-api/internal/domain/entity"
	gameenum "github.com/qilin/crm-api/internal/domain/enum/game"
	"github.com/qilin/crm-api/internal/domain/enum/game_rating"
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
	Trailer      *string        `json:"trailer"`
	Platforms    *[]string      `json:"platforms"`
	Developers   *[]uint        `json:"developers"`
	Features     *[]uint        `json:"features"`
	Genres       *[]uint        `json:"genres"`
	Publishers   *[]uint        `json:"publishers"`
	Tags         *[]uint        `json:"tags"`
	Media        *[]uint        `json:"media"`
	ReleaseDate  *time.Time     `json:"release_date"`
	SocialLinks  []socialLink   `json:"social_links"`
	Localization []localization `json:"localization"`
	Raiting      *[]rating      `json:"raiting"`
	Review       *[]review      `json:"review"`
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
			Trailer:     req.Trailer,
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

	if req.Raiting != nil {
		ratings := make([]service.RatingData, len(*req.Raiting))
		for i, r := range *req.Raiting {
			ratings[i] = service.RatingData{
				Agency:              game_rating.NewAgencyByString(r.Agency),
				Rating:              game_rating.NewRatingByString(r.Agency, r.Rating),
				DisplayOnlineNotice: r.DisplayOnlineNotice,
				ShowAgeRestrict:     r.ShowAgeRestrict,
			}
		}
		data.CommonGameData.Ratings = &ratings
	}

	if req.Review != nil {
		reviews := make([]service.ReviewData, len(*req.Review))
		for i, r := range *req.Review {
			reviews[i] = service.ReviewData{
				PressName: r.PressName,
				Link:      r.Link,
				Score:     r.Score,
				Quote:     r.Quote,
			}
		}
		data.CommonGameData.Reviews = &reviews
	}

	return data, nil
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
