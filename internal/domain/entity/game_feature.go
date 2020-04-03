package entity

type GameFeature struct {
	ID        uint
	GameID    uint
	FeatureID uint
}

type GameFeatureArray struct {
	ids        []uint
	featureIDs []uint
	items      []GameFeature
}

func NewGameFeatureArray(items []GameFeature) *GameFeatureArray {
	a := &GameFeatureArray{
		items: items,
	}
	a.refresh()
	return a
}

func (a *GameFeatureArray) refresh() {
	ids := make([]uint, len(a.items))
	featureIDs := make([]uint, len(a.items))

	for i := range a.items {
		ids[i] = a.items[i].ID
		featureIDs[i] = a.items[i].FeatureID
	}

	a.ids = ids
	a.featureIDs = featureIDs
}

func (a *GameFeatureArray) Size() int {
	return len(a.items)
}

func (a *GameFeatureArray) IDs() []uint {
	items := make([]uint, len(a.items))
	copy(items, a.ids)
	return items
}

func (a *GameFeatureArray) FeatureIDs() []uint {
	items := make([]uint, len(a.items))
	copy(items, a.featureIDs)
	return items
}
