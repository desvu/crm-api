package entity

type GameRevisionTag struct {
	ID             uint
	GameRevisionID uint
	TagID          uint
}

type GameRevisionTagArray struct {
	ids    []uint
	tagIDs []uint
	items  []GameRevisionTag
}

func NewGameRevisionTagArray(items []GameRevisionTag) *GameRevisionTagArray {
	a := &GameRevisionTagArray{
		items: items,
	}
	a.refresh()
	return a
}

func (a *GameRevisionTagArray) refresh() {
	ids := make([]uint, len(a.items))
	tagIDs := make([]uint, len(a.items))

	for i := range a.items {
		ids[i] = a.items[i].ID
		tagIDs[i] = a.items[i].TagID
	}

	a.ids = ids
	a.tagIDs = tagIDs
}

func (a *GameRevisionTagArray) Size() int {
	return len(a.items)
}

func (a *GameRevisionTagArray) IDs() []uint {
	items := make([]uint, len(a.items))
	copy(items, a.ids)
	return items
}

func (a *GameRevisionTagArray) TagIDs() []uint {
	items := make([]uint, len(a.items))
	copy(items, a.tagIDs)
	return items
}
