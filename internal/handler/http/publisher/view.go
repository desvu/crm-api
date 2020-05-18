package publisher

import (
	"github.com/qilin/crm-api/internal/domain/entity"
)

//swagger:model PublisherList
type publisherList []publisher

//swagger:model publisher
type publisher struct {
	// example: 12
	ID uint `json:"id"`

	// example: Acme Corp
	Name string `json:"name"`
}

func (h *Handler) view(t entity.Publisher) publisher {
	return publisher{
		ID:   t.ID,
		Name: t.Name,
	}
}

func (h *Handler) viewList(items []entity.Publisher) publisherList {
	var res = make([]publisher, len(items))
	for i := range items {
		res[i] = h.view(items[i])
	}
	return res
}
