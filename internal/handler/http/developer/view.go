package developer

import (
	"github.com/qilin/crm-api/internal/domain/entity"
)

//swagger:model DeveloperList
type developerList []developer

//swagger:model genre
type developer struct {
	// example: 12
	ID uint `json:"id"`

	// example: Acme Corp
	Name string `json:"name"`
}

func (h *Handler) view(t entity.Developer) developer {
	return developer{
		ID:   t.ID,
		Name: t.Name,
	}
}

func (h *Handler) viewList(items []entity.Developer) developerList {
	var res = make([]developer, len(items))
	for i := range items {
		res[i] = h.view(items[i])
	}
	return res
}
