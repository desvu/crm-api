package genre

import (
	"github.com/qilin/crm-api/internal/domain/entity"
)

//swagger:model GenreList
type genreList []genre

//swagger:model genre
type genre struct {
	// read-only: true
	// example: 12
	ID uint `json:"id"`

	// example: Strategy
	// read-only: true
	Name string `json:"name"`
}

func (h *Handler) view(t entity.Genre) genre {
	return genre{
		ID:   t.ID,
		Name: t.Name,
	}
}

func (h *Handler) viewList(tags []entity.Genre) genreList {
	var res = make([]genre, len(tags))
	for i := range tags {
		res[i] = h.view(tags[i])
	}
	return res
}
