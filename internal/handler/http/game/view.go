package game

import "github.com/qilin/crm-api/internal/domain/entity"

type view struct {
	Game game `json:"game"`
}

type game struct {
	ID       string   `json:"id"`
	Title    string   `json:"title"`
	Type     string   `json:"type"`
	Slug     string   `json:"slug"`
	Revision revision `json:"revision"`
}

type revision struct {
	ID          uint         `json:"id"`
	Status      string       `json:"status"`
	Summary     string       `json:"summary,omitempty"`
	Description string       `json:"description,omitempty"`
	License     string       `json:"license,omitempty"`
	Media       []media      `json:"media,omitempty"`
	SocialLinks []socialLink `json:"social_links,omitempty"`
}

type socialLink struct {
	URL string `json:"url"`
}

type media struct {
	ID   uint   `json:"id"`
	Type string `json:"type"`
	URL  string `json:"url"`
}

func (h Handler) view(i *entity.GameEx) view {
	var v = view{Game: game{
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
			SocialLinks: convertEntitySocialLinksToSocialLinks(i.Revision.SocialLinks),
		},
	}}

	if len(i.Revision.Media) > 0 {
		for _, m := range i.Revision.Media {
			v.Game.Revision.Media = append(v.Game.Revision.Media, media{
				ID:   m.ID,
				Type: m.Type.String(),
				URL:  h.URLBuilder.BuildGameMedia(&m),
			})
		}
	}

	return v
}
