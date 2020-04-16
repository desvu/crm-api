package entity

type GameRevisionFeature struct {
	ID             uint
	GameRevisionID uint
	FeatureID      uint
}

type GameRevisionFeatureArray struct {
	ids        []uint
	featureIDs []uint
	items      []GameRevisionFeature
}

func NewGameRevisionFeatureArray(items []GameRevisionFeature) *GameRevisionFeatureArray {
	a := &GameRevisionFeatureArray{
		items: items,
	}
	a.refresh()
	return a
}

func (a *GameRevisionFeatureArray) refresh() {
	ids := make([]uint, len(a.items))
	featureIDs := make([]uint, len(a.items))

	for i := range a.items {
		ids[i] = a.items[i].ID
		featureIDs[i] = a.items[i].FeatureID
	}

	a.ids = ids
	a.featureIDs = featureIDs
}

func (a *GameRevisionFeatureArray) Size() int {
	return len(a.items)
}

func (a *GameRevisionFeatureArray) IDs() []uint {
	items := make([]uint, len(a.items))
	copy(items, a.ids)
	return items
}

func (a *GameRevisionFeatureArray) FeatureIDs() []uint {
	items := make([]uint, len(a.items))
	copy(items, a.featureIDs)
	return items
}
