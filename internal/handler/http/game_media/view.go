package game_media

import "github.com/qilin/crm-api/internal/domain/entity"

type view struct {
	Media media `json:"game_media"`
}

type media struct {
	ID         uint   `json:"id"`
	Type       string `json:"type"`
	GameID     string `json:"game_id"`
	IsUploaded bool   `json:"is_uploaded"`
	URL        string `json:"url"`
}

func (h Handler) view(i *entity.GameMedia) view {
	return view{Media: media{
		ID:         i.ID,
		Type:       i.Type.String(),
		GameID:     i.GameID,
		IsUploaded: i.IsUploaded,
		URL:        h.URLBuilder.BuildGameMedia(i),
	}}
}
