package entity

type GameTag struct {
	ID     uint
	GameID uint
	TagID  uint
}

type GameTagArray struct {
	ids   []uint
	items []GameTag
}

func NewGameTagArray(gameTags []GameTag) *GameTagArray {
	a := &GameTagArray{
		items: gameTags,
	}
	a.refresh()
	return a
}

func (a *GameTagArray) refresh() {
	ids := make([]uint, len(a.items))

	for i := range a.items {
		ids[i] = a.items[i].ID
	}

	a.ids = ids
}

func (a *GameTagArray) Size() int {
	return len(a.items)
}

func (a *GameTagArray) IDs() []uint {
	items := make([]uint, len(a.items))
	copy(items, a.ids)
	return items
}
