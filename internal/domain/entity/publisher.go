package entity

type Publisher struct {
	ID   uint
	Name string
}

type PublisherArray struct {
	ids   []uint
	items []Publisher
}

func NewPublisherArray(items []Publisher) *PublisherArray {
	a := &PublisherArray{
		items: items,
	}
	a.refresh()
	return a
}

func (a *PublisherArray) refresh() {
	ids := make([]uint, len(a.items))

	for i := range a.items {
		ids[i] = a.items[i].ID
	}

	a.ids = ids
}

func (a *PublisherArray) Size() int {
	return len(a.items)
}

func (a *PublisherArray) IDs() []uint {
	items := make([]uint, len(a.items))
	copy(items, a.ids)
	return items
}
