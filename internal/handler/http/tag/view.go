package tag

import (
	"github.com/qilin/crm-api/internal/domain/entity"
)

//swagger:model TagList
type tagList []tag

//swagger:model tag
type tag struct {
	// read-only: true
	// example: 12
	ID uint `json:"id"`

	// example: RPG
	// read-only: true
	Name string `json:"name"`
}

func (h *Handler) view(t entity.Tag) tag {
	return tag{
		ID:   t.ID,
		Name: t.Name,
	}
}

func (h *Handler) viewList(tags []entity.Tag) tagList {
	var res = make([]tag, len(tags))
	for i := range tags {
		res[i] = h.view(tags[i])
	}
	return res
}
