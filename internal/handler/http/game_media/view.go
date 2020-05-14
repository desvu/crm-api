package game_media

import "github.com/qilin/crm-api/internal/domain/entity"

//swagger:model Media
type media struct {
	// example: 32
	ID uint `json:"id"`

	// enum: wideSlider,vertical,horizontal,horizontalSmall,largeSingle,catalog,screenshot,description
	Type string `json:"type"`

	// example: b6fa8b92-8d5b-42e5-a7e8-3e5dabb2ca51
	GameID string `json:"game_id"`

	//example: true
	IsUploaded bool `json:"is_uploaded"`

	// example: https://sdn.qilin.super.com/game/b6fa8b92-8d5b-42e5-a7e8-3e5dabb2ca51/media/b6fa8b92-8d5b-42e5-a7e8-3e5dabb2ca51.png
	URL string `json:"url"`
}

func (h Handler) view(i *entity.GameMedia) media {
	return media{
		ID:         i.ID,
		Type:       i.Type.String(),
		GameID:     i.GameID,
		IsUploaded: i.IsUploaded,
		URL:        h.URLBuilder.BuildGameMedia(i),
	}
}
