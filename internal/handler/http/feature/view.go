package feature

import (
	"github.com/qilin/crm-api/internal/domain/entity"
)

//swagger:model FeatureList
type featureList []feature

//swagger:model genre
type feature struct {
	// example: 12
	ID uint `json:"id"`

	// example: ControllerFullSupport
	Name string `json:"name"`

	// example: gamepad
	Icon string `json:"icon"`
}

func (h *Handler) view(t entity.Feature) feature {
	return feature{
		ID:   t.ID,
		Name: t.Name,
		Icon: t.Icon.String(),
	}
}

func (h *Handler) viewList(tags []entity.Feature) featureList {
	var res = make([]feature, len(tags))
	for i := range tags {
		res[i] = h.view(tags[i])
	}
	return res
}
