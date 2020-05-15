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
	// required: true
	// example: 11002485-cb51-4b29-8423-cba43f29f143
	ID string `param:"game_id"`
}

//swagger:parameters reqUpsert
type reqUpsertWrapper struct {
	// in: body
	Data reqUpsert
}
type reqUpsert struct {
	// example: 11002485-cb51-4b29-8423-cba43f29f143
	ID *string `json:"id"`

	// example: Ash of Gods
	Title *string `json:"title"`

	// example: desktop
	Type *string `json:"type"`

	// example: ash-of-gods
	Slug *string `json:"slug"`

	// example: Ash of Gods: Redemption is a turn-based RPG that combines tactical combat, CCG elements, and a constantly evolving story in which no one is safe from death, including the main characters.
	Summary *string `json:"summary"`

	// example: Ash of Gods is the story of three separate protagonists rising in response to a centuries-old menace once thought to be mere folklore. Captain Thorn Brenin, the bodyguard Lo Pheng, the scribe Hopper Rouley, and many others, do not yet know that the reapers have returned and intend to drown the world in blood so that they may awaken the sleeping gods.
	Description *string `json:"description"`
	License     *string `json:"license"`

	// example: https://www.youtube.com/watch?v=k_j0fw8jh8M
	Trailer *string `json:"trailer"`

	// example: 360
	PlayTime *uint `json:"play_time"`

	// example: [windows, macOS]
	Platforms *[]gameenum.Platform `json:"platforms"`

	// example: [32]
	Developers *[]uint `json:"developers"`

	// example: [14, 44, 67]
	Features *[]uint `json:"features"`

	// example: [53, 23, 1]
	Genres *[]uint `json:"genres"`

	// example: [1]
	Publishers *[]uint `json:"publishers"`

	// example: []
	Tags *[]uint `json:"tags"`

	// example: [2, 5, 8, 9]
	Media        *[]uint         `json:"media"`
	ReleaseDate  *time.Time      `json:"release_date"`
	SocialLinks  *[]socialLink   `json:"social_links"`
	Localization *[]localization `json:"localization"`
	Rating       *[]rating       `json:"rating"`
	Review       *[]review       `json:"review"`
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
			PlayTime:    req.PlayTime,
			Tags:        req.Tags,
			Developers:  req.Developers,
			Publishers:  req.Publishers,
			Features:    req.Features,
			Genres:      req.Genres,
			Media:       req.Media,
			Platforms:   nil, // todo ?
			ReleaseDate: req.ReleaseDate,
			SocialLinks: convertSocialLinksToServiceSocialLinks(req.SocialLinks),
		},
	}

	if req.Localization != nil {
		localizations := make([]service.LocalizationData, len(*req.Localization))
		for i, l := range *req.Localization {
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

func convertSocialLinksToServiceSocialLinks(links *[]socialLink) *[]service.SocialLink {
	if links == nil {
		return nil
	}
	list := make([]service.SocialLink, len(*links))
	for i, item := range *links {
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

//swagger:parameters reqGetByFilter
type reqGetByFilter struct {
	// in: query
	// example: 30
	Limit int `query:"limit"`

	// in: body
	// example: 20
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
