package game

import "github.com/qilin/crm-api/internal/domain/entity"

//swagger:model Game
type game struct {
	// read-only: true
	// example: b6fa8b92-8d5b-42e5-a7e8-3e5dabb2ca51
	ID string `json:"id"`

	// example: Ash of Gods
	Title string `json:"title"`

	// example: desktop
	Type string `json:"type"`

	// example: ash-of-gods
	Slug string `json:"slug"`

	Revision revision `json:"revision"`
}

type revision struct {
	// read-only: true
	// example: 43
	ID uint `json:"id"`

	// example: published
	Status string `json:"status"`

	// example: Summary game
	Summary string `json:"summary,omitempty"`

	// example: Description game
	Description  string         `json:"description,omitempty"`
	License      string         `json:"license,omitempty"`
	Trailer      string         `json:"trailer,omitempty"`
	Media        []media        `json:"media,omitempty"`
	SocialLinks  []socialLink   `json:"social_links,omitempty"`
	Localization []localization `json:"localization,omitempty"`
	Rating       []rating       `json:"rating,omitempty"`
}

type socialLink struct {
	URL string `json:"url"`
}

type media struct {
	ID   uint   `json:"id"`
	Type string `json:"type"`
	URL  string `json:"url"`
}

type localization struct {
	Language  string `json:"language"`
	Interface bool   `json:"interface"`
	Audio     bool   `json:"audio"`
	Subtitles bool   `json:"subtitles"`
}

type rating struct {
	Agency              string `json:"agency"`
	Rating              string `json:"rating"`
	DisplayOnlineNotice bool   `json:"display_online_notice"`
	ShowAgeRestrict     bool   `json:"show_age_restrict"`
}

func (h Handler) view(i *entity.GameEx) game {
	var v = game{
		ID:    i.ID,
		Title: i.Title,
		Type:  i.Type.String(),
		Slug:  i.Slug,
		Revision: revision{
			ID:          i.Revision.ID,
			Status:      i.Revision.Status.String(),
			Summary:     i.Revision.Summary,
			Description: i.Revision.Description,
			License:     i.Revision.License,
			Trailer:     i.Revision.Trailer,
			SocialLinks: convertEntitySocialLinksToSocialLinks(i.Revision.SocialLinks),
		},
	}

	if len(i.Revision.Media) > 0 {
		for _, m := range i.Revision.Media {
			v.Revision.Media = append(v.Revision.Media, media{
				ID:   m.ID,
				Type: m.Type.String(),
				URL:  h.URLBuilder.BuildGameMedia(&m),
			})
		}
	}

	if len(i.Revision.Localization) > 0 {
		for _, l := range i.Revision.Localization {
			v.Revision.Localization = append(v.Revision.Localization, localization{
				Language:  l.Language,
				Interface: l.Interface,
				Audio:     l.Audio,
				Subtitles: l.Subtitles,
			})
		}
	}

	if len(i.Revision.Rating) > 0 {
		for _, r := range i.Revision.Rating {
			v.Revision.Rating = append(v.Revision.Rating, rating{
				Agency:              r.Agency.String(),
				Rating:              r.Rating.String(),
				DisplayOnlineNotice: r.DisplayOnlineNotice,
				ShowAgeRestrict:     r.ShowAgeRestrict,
			})
		}
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
