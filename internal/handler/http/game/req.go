package game

import (
	"time"

	"github.com/qilin/crm-api/internal/domain/enum/game_social_link"

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
	// ID is required on revision update
	//
	// example: 11002485-cb51-4b29-8423-cba43f29f143
	ID *string `json:"id"`

	// example: Ash of Gods
	// required: true
	Title *string `json:"title"`

	// example: desktop
	// required: true
	Type *string `json:"type"`

	// example: ash-of-gods
	// required: true
	Slug *string `json:"slug"`

	// example: Ash of Gods: Redemption is a turn-based RPG that combines tactical combat, CCG elements, and a constantly evolving story in which no one is safe from death, including the main characters.
	Summary *string `json:"summary"`

	// example: Ash of Gods is the story of three separate protagonists rising in response to a centuries-old menace once thought to be mere folklore. Captain Thorn Brenin, the bodyguard Lo Pheng, the scribe Hopper Rouley, and many others, do not yet know that the reapers have returned and intend to drown the world in blood so that they may awaken the sleeping gods.
	Description *string `json:"description"`

	//
	License *string `json:"license"`

	// example: https://www.youtube.com/watch?v=k_j0fw8jh8M
	Trailer *string `json:"trailer"`

	// example: 360
	PlayTime *uint `json:"play_time"`

	// example: ["windows", "macOS"]
	Platforms *[]string `json:"platforms"`

	// example: [4]
	Developers *[]uint `json:"developers"`

	// example: [1, 3]
	Features *[]uint `json:"features"`

	// example: [2, 1]
	Genres *[]uint `json:"genres"`

	// example: [4,6,7]
	Publishers *[]uint `json:"publishers"`

	// example: [3, 2, 6, 4]
	Tags *[]uint `json:"tags"`

	// example: [2, 5, 8, 9]
	Media *[]uint `json:"media"`

	// example: 2020-01-02T00:00:00Z
	ReleaseDate *time.Time `json:"release_date"`

	// example: [{"type": "facebook", "url": "https://www.facebook.com/AshOfGods/"},{"type": "reddit", "url":"https://www.reddit.com/r/ashofgods/"}]
	SocialLinks *[]socialLink `json:"social_links"`

	// example: [{"language": "eng", "interface": true, "audio": true, "subtitles": true}]
	Localization *[]localization `json:"localization"`

	// example: [{"agency": "CERO", "rating": "A", "display_online_notice": true, "show_age_restrict": true}]
	Rating *[]rating `json:"rating"`

	// example: [{ "press_name": "IGN", "link": "https://www.ign.com/articles/2013/09/16/grand-theft-auto-v-review", "score": 90, "quote": "Grand Theft Auto V is not only a preposterously enjoyable video game, but also an intelligent and sharp-tongued satire of contemporary America." }]
	Review *[]review `json:"review"`

	// example: [{"platform": "windows", "minimal": {"os": "Windows 7/8 64-bit", "cpu": "i5", "gpu": "GTC 1050", "disk_space": 6500, "ram": 6000}, "recommended": {"os": "Windows 8/10 64-bit", "cpu": "i7", "gpu": "GTC 1080", "disk_space": 6500, "ram": 8000}}]
	SystemRequirements *[]systemRequirements `json:"system_requirements"`
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
			ReleaseDate: req.ReleaseDate,
			SocialLinks: convertSocialLinksToServiceSocialLinks(req.SocialLinks),
		},
	}

	if req.Platforms != nil {
		platforms := gameenum.NewPlatformArrayByStrings((*req.Platforms)...)
		data.CommonGameData.Platforms = &platforms
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

	if req.SystemRequirements != nil {
		requirements := make([]service.SystemRequirements, len(*req.SystemRequirements))
		for i, r := range *req.SystemRequirements {
			requirements[i] = service.SystemRequirements{
				Platform: gameenum.NewPlatformByString(r.Platform),
			}
			if r.Minimal != nil {
				requirements[i].Minimal = &service.RequirementsSet{
					CPU:       r.Minimal.CPU,
					GPU:       r.Minimal.GPU,
					DiskSpace: r.Minimal.DiskSpace,
					RAM:       r.Minimal.RAM,
				}
			}
			if r.Recommended != nil {
				requirements[i].Recommended = &service.RequirementsSet{
					CPU:       r.Recommended.CPU,
					GPU:       r.Recommended.GPU,
					DiskSpace: r.Recommended.DiskSpace,
					RAM:       r.Recommended.RAM,
				}
			}
		}
		data.CommonGameData.SystemRequirements = &requirements
	}

	return data, nil
}

func convertSocialLinksToServiceSocialLinks(links *[]socialLink) *[]service.SocialLink {
	if links == nil {
		return nil
	}
	list := make([]service.SocialLink, len(*links))
	for i, item := range *links {
		list[i] = service.SocialLink{Type: game_social_link.NewTypeByString(item.Type), URL: item.URL}
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

	// in: query
	// example: 20
	Offset int `query:"offset"`
}

func convertGetByFilterRequest(c echo.Context) (*service.GetByFilterGameData, error) {
	req := new(reqGetByFilter)
	if err := c.Bind(req); err != nil {
		return nil, err
	}

	return &service.GetByFilterGameData{
		Limit:  req.Limit,
		Offset: req.Offset,
	}, nil
}
