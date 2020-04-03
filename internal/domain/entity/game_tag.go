package entity

type GameTag struct {
	ID     uint
	GameID uint
	TagID  uint
}

type GameTagArray struct {
	ids    []uint
	tagIDs []uint
	items  []GameTag
}

func NewGameTagArray(items []GameTag) *GameTagArray {
	a := &GameTagArray{
		items: items,
	}
	a.refresh()
	return a
}

func (a *GameTagArray) refresh() {
	ids := make([]uint, len(a.items))
	tagIDs := make([]uint, len(a.items))

	for i := range a.items {
		ids[i] = a.items[i].ID
		tagIDs[i] = a.items[i].TagID
	}

	a.ids = ids
	a.tagIDs = tagIDs
}

func (a *GameTagArray) Size() int {
	return len(a.items)
}

func (a *GameTagArray) IDs() []uint {
	items := make([]uint, len(a.items))
	copy(items, a.ids)
	return items
}

func (a *GameTagArray) TagIDs() []uint {
	items := make([]uint, len(a.items))
	copy(items, a.tagIDs)
	return items
}
