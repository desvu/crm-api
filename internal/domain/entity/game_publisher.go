package entity

type GamePublisher struct {
	ID          uint
	GameID      uint
	PublisherID uint
}

type GamePublisherArray struct {
	ids          []uint
	publisherIDs []uint
	items        []GamePublisher
}

func NewGamePublisherArray(items []GamePublisher) *GamePublisherArray {
	a := &GamePublisherArray{
		items: items,
	}
	a.refresh()
	return a
}

func (a *GamePublisherArray) refresh() {
	ids := make([]uint, len(a.items))
	publisherIDs := make([]uint, len(a.items))

	for i := range a.items {
		ids[i] = a.items[i].ID
		publisherIDs[i] = a.items[i].PublisherID
	}

	a.ids = ids
	a.publisherIDs = publisherIDs
}

func (a *GamePublisherArray) Size() int {
	return len(a.items)
}

func (a *GamePublisherArray) IDs() []uint {
	items := make([]uint, len(a.items))
	copy(items, a.ids)
	return items
}

func (a *GamePublisherArray) PublisherIDs() []uint {
	items := make([]uint, len(a.items))
	copy(items, a.publisherIDs)
	return items
}
