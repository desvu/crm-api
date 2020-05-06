package entity

type GameRevisionMedia struct {
	ID         uint
	RevisionID uint
	MediaID    uint
}

type GameRevisionMediaArray struct {
	ids      []uint
	mediaIDs []uint
	items    []GameRevisionMedia
}

func NewGameRevisionMediaArray(items []GameRevisionMedia) *GameRevisionMediaArray {
	a := &GameRevisionMediaArray{
		items: items,
	}
	a.refresh()
	return a
}

func (a *GameRevisionMediaArray) refresh() {
	ids := make([]uint, len(a.items))
	mediaIDs := make([]uint, len(a.items))

	for i := range a.items {
		ids[i] = a.items[i].ID
		mediaIDs[i] = a.items[i].MediaID
	}

	a.ids = ids
	a.mediaIDs = mediaIDs
}

func (a *GameRevisionMediaArray) Size() int {
	return len(a.items)
}

func (a *GameRevisionMediaArray) IDs() []uint {
	items := make([]uint, len(a.items))
	copy(items, a.ids)
	return items
}

func (a *GameRevisionMediaArray) MediaIDs() []uint {
	items := make([]uint, len(a.items))
	copy(items, a.mediaIDs)
	return items
}
