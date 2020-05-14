package game

import (
	"time"

	"github.com/labstack/echo/v4"
	"github.com/qilin/crm-api/internal/domain/entity"
	gameenum "github.com/qilin/crm-api/internal/domain/enum/game"
	"github.com/qilin/crm-api/internal/domain/enum/game_rating"
	"github.com/qilin/crm-api/internal/domain/service"
)

//swagger:parameters reqGetByID reqPublish
type reqByID struct {
	// in: path
	// example: 11002485-cb51-4b29-8423-cba43f29f143
	ID string `param:"game_id"`
}

//swagger:parameters reqUpsert
type reqUpsert struct {
	// in: body
	// example: 11002485-cb51-4b29-8423-cba43f29f143
	ID *string `json:"id"`
	// in: body
	Title *string `json:"title"`
	// in: body
	Type *string `json:"type"`
	// in: body
	Slug *string `json:"slug"`
	// in: body
	Summary *string `json:"summary"`
	// in: body
	Description *string `json:"description"`
	// in: body
	License *string `json:"license"`
	// in: body
	Trailer *string `json:"trailer"`
	// in: body
	Platforms *[]string `json:"platforms"`
	// in: body
	Developers *[]uint `json:"developers"`
	// in: body
	Features *[]uint `json:"features"`
	// in: body
	Genres *[]uint `json:"genres"`
	// in: body
	Publishers *[]uint `json:"publishers"`
	// in: body
	Tags *[]uint `json:"tags"`
	// in: body
	Media *[]uint `json:"media"`
	// in: body
	ReleaseDate *time.Time `json:"release_date"`
	// in: body
	SocialLinks []socialLink `json:"social_links"`
	// in: body
	Localization []localization `json:"localization"`
	// in: body
	Rating *[]rating `json:"rating"`
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

	if req.Rating != nil {
		ratings := make([]service.RatingData, len(*req.Rating))
		for i, r := range *req.Rating {
			ratings[i] = service.RatingData{
				Agency:              game_rating.NewAgencyByString(r.Agency),
				Rating:              game_rating.NewRatingByString(r.Agency, r.Rating),
				DisplayOnlineNotice: r.DisplayOnlineNotice,
				ShowAgeRestrict:     r.ShowAgeRestrict,
			}
		}
		data.CommonGameData.Ratings = &ratings
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

type reqGetByFilter struct {
	Limit  int `query:"limit"`
	Offset int `query:"offset"`
}

func convertGetByFilterRequest(c echo.Context) (*service.GetByFilterGameDate, error) {
	req := new(reqGetByFilter)
	if err := c.Bind(req); err != nil {
		return nil, err
	}

	return &service.GetByFilterGameDate{
		Limit:  req.Limit,
		Offset: req.Offset,
	}, nil
}
