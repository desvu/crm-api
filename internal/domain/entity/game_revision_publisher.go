package entity

type GameRevisionPublisher struct {
	ID             uint
	GameRevisionID uint
	PublisherID    uint
}

type GameRevisionPublisherArray struct {
	ids          []uint
	publisherIDs []uint
	items        []GameRevisionPublisher
}

func NewGameRevisionPublisherArray(items []GameRevisionPublisher) *GameRevisionPublisherArray {
	a := &GameRevisionPublisherArray{
		items: items,
	}
	a.refresh()
	return a
}

func (a *GameRevisionPublisherArray) refresh() {
	ids := make([]uint, len(a.items))
	publisherIDs := make([]uint, len(a.items))

	for i := range a.items {
		ids[i] = a.items[i].ID
		publisherIDs[i] = a.items[i].PublisherID
	}

	a.ids = ids
	a.publisherIDs = publisherIDs
}

func (a *GameRevisionPublisherArray) Size() int {
	return len(a.items)
}

func (a *GameRevisionPublisherArray) IDs() []uint {
	items := make([]uint, len(a.items))
	copy(items, a.ids)
	return items
}

func (a *GameRevisionPublisherArray) PublisherIDs() []uint {
	items := make([]uint, len(a.items))
	copy(items, a.publisherIDs)
	return items
}
