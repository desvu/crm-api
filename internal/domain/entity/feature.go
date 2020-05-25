package entity

import "github.com/qilin/crm-api/internal/domain/enum/game"

type Feature struct {
	ID   uint
	Name string
	Icon game.Icon
}

type FeatureArray struct {
	ids   []uint
	items []Feature
}

func NewFeatureArray(items []Feature) *FeatureArray {
	a := &FeatureArray{
		items: items,
	}
	a.refresh()
	return a
}

func (a *FeatureArray) refresh() {
	ids := make([]uint, len(a.items))

	for i := range a.items {
		ids[i] = a.items[i].ID
	}

	a.ids = ids
}

func (a *FeatureArray) Size() int {
	return len(a.items)
}

func (a *FeatureArray) IDs() []uint {
	items := make([]uint, len(a.items))
	copy(items, a.ids)
	return items
}
