package game

import "github.com/qilin/crm-api/internal/domain/entity"

type view struct {
	ID       string   `json:"id"`
	Title    string   `json:"title"`
	Type     string   `json:"type"`
	Slug     string   `json:"slug"`
	Revision revision `json:"revision"`
}

type revision struct {
	ID          uint    `json:"id"`
	Status      string  `json:"status"`
	Summary     string  `json:"summary,omitempty"`
	Description string  `json:"description,omitempty"`
	License     string  `json:"license,omitempty"`
	Trailer     string  `json:"trailer,omitempty"`
	Media       []media `json:"media,omitempty"`
    SocialLinks []socialLink `json:"social_links,omitempty"`
    Localization []localization `json:"localization,omitempty"`
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

func (h Handler) view(i *entity.GameEx) view {
	var v = view{
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

	return v
}
