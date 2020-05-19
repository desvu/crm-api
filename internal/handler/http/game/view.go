package game

import (
	"github.com/qilin/crm-api/internal/domain/entity"
	gameenum "github.com/qilin/crm-api/internal/domain/enum/game"
	"github.com/qilin/crm-api/internal/domain/enum/game_media"
)

//swagger:model Game
type game struct {
	// read-only: true
	// example: b6fa8b92-8d5b-42e5-a7e8-3e5dabb2ca51
	ID string `json:"id"`

	// example: Ash of Gods
	Title string `json:"title"`

	// example: desktop
	Type gameenum.Type `json:"type"`

	// example: ash-of-gods
	Slug string `json:"slug"`

	Revision revision `json:"revision"`
}

type revision struct {
	// read-only: true
	// example: 43
	ID uint `json:"id"`

	// example: published
	// enum: draft,publishing,published
	Status string `json:"status"`

	// example: Summary game
	Summary string `json:"summary,omitempty"`

	// example: Description game
	Description        string               `json:"description,omitempty"`
	License            string               `json:"license,omitempty"`
	Trailer            string               `json:"trailer,omitempty"`
	PlayTime           uint                 `json:"play_time,omitempty"`
	Platforms          []string             `json:"platforms"`
	Media              []media              `json:"media,omitempty"`
	SocialLinks        []socialLink         `json:"social_links,omitempty"`
	Localization       []localization       `json:"localization,omitempty"`
	Rating             []rating             `json:"rating,omitempty"`
	Review             []review             `json:"review,omitempty"`
	SystemRequirements []systemRequirements `json:"system_requirements,omitempty"`
}

type socialLink struct {
	// required: true
	URL string `json:"url"`
}

type media struct {
	ID   uint            `json:"id"`
	Type game_media.Type `json:"type"`
	URL  string          `json:"url"`
}

type localization struct {
	// required: true
	Language string `json:"language"`
	// required: true
	Interface bool `json:"interface"`
	// required: true
	Audio bool `json:"audio"`
	// required: true
	Subtitles bool `json:"subtitles"`
}

type rating struct {
	// required: true
	Agency string `json:"agency"`
	// required: true
	Rating string `json:"rating"`
	// required: true
	DisplayOnlineNotice bool `json:"display_online_notice"`
	// required: true
	ShowAgeRestrict bool `json:"show_age_restrict"`
}

type review struct {
	// required: true
	PressName string `json:"press_name"`
	// required: true
	Link string `json:"link"`
	// required: true
	Score string `json:"score"`
	// required: true
	Quote string `json:"quote"`
}

type systemRequirements struct {
	// required: true
	// example: windows
	Platform string `json:"platform"`
	// example: {"cpu": "i5", "gpu": "GTC 1050", "disk_space": 6500, "ram": 6000}
	Minimal *requirementsSet `json:"minimal,omitempty"`
	// example: {"cpu": "i7", "gpu": "GTC 1080", "disk_space": 6500, "ram": 8000}
	Recommended *requirementsSet `json:"recommended,omitempty"`
}

type requirementsSet struct {
	CPU       string `json:"cpu"`
	GPU       string `json:"gpu"`
	DiskSpace uint   `json:"disk_space"`
	RAM       uint   `json:"ram"`
}

func (h Handler) view(i *entity.GameEx) game {
	var v = game{
		ID:    i.ID,
		Title: i.Title,
		Type:  i.Type,
		Slug:  i.Slug,
		Revision: revision{
			ID:          i.Revision.ID,
			Status:      i.Revision.Status.String(),
			Summary:     i.Revision.Summary,
			Description: i.Revision.Description,
			License:     i.Revision.License,
			Trailer:     i.Revision.Trailer,
			PlayTime:    i.Revision.PlayTime,
			Platforms:   i.Revision.Platforms.Strings(),
			SocialLinks: convertEntitySocialLinksToSocialLinks(i.Revision.SocialLinks),
		},
	}

	for _, m := range i.Revision.Media {
		v.Revision.Media = append(v.Revision.Media, media{
			ID:   m.ID,
			Type: m.Type,
			URL:  h.URLBuilder.BuildGameMedia(&m),
		})
	}

	for _, l := range i.Revision.Localization {
		v.Revision.Localization = append(v.Revision.Localization, localization{
			Language:  l.Language,
			Interface: l.Interface,
			Audio:     l.Audio,
			Subtitles: l.Subtitles,
		})
	}

	for _, r := range i.Revision.Rating {
		v.Revision.Rating = append(v.Revision.Rating, rating{
			Agency:              r.Agency.String(),
			Rating:              r.Rating.String(),
			DisplayOnlineNotice: r.DisplayOnlineNotice,
			ShowAgeRestrict:     r.ShowAgeRestrict,
		})
	}

	for _, r := range i.Revision.Review {
		v.Revision.Review = append(v.Revision.Review, review{
			PressName: r.PressName,
			Link:      r.Link,
			Score:     r.Score,
			Quote:     r.Quote,
		})
	}

	for _, r := range i.Revision.SystemRequirements {
		set := systemRequirements{
			Platform: r.Platform.String(),
		}
		if r.Minimal != nil {
			set.Minimal = &requirementsSet{
				CPU:       r.Minimal.CPU,
				GPU:       r.Minimal.GPU,
				DiskSpace: r.Minimal.DiskSpace,
				RAM:       r.Minimal.RAM,
			}
		}
		if r.Recommended != nil {
			set.Recommended = &requirementsSet{
				CPU:       r.Recommended.CPU,
				GPU:       r.Recommended.GPU,
				DiskSpace: r.Recommended.DiskSpace,
				RAM:       r.Recommended.RAM,
			}
		}
		v.Revision.SystemRequirements = append(v.Revision.SystemRequirements, set)
	}

	return v
}

type pagination struct {
	Total int `json:"total"`
}

//swagger:model GameList
type gameList struct {
	Games      []game     `json:"games"`
	Pagination pagination `json:"pagination"`
}

func (h Handler) viewArray(items []entity.GameEx) gameList {
	var games = make([]game, len(items))
	for i := range items {
		games[i] = h.view(&items[i])
	}

	return gameList{
		Games: games,
		Pagination: pagination{
			Total: 0,
		},
	}
}
