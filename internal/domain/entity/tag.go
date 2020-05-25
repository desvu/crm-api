package entity

type Tag struct {
	ID   uint
	Name string
}

type TagArray struct {
	ids   []uint
	items []Tag
}

func NewTagArray(items []Tag) *TagArray {
	a := &TagArray{
		items: items,
	}
	a.refresh()
	return a
}

func (a *TagArray) refresh() {
	ids := make([]uint, len(a.items))

	for i := range a.items {
		ids[i] = a.items[i].ID
	}

	a.ids = ids
}

func (a *TagArray) Size() int {
	return len(a.items)
}

func (a *TagArray) IDs() []uint {
	items := make([]uint, len(a.items))
	copy(items, a.ids)
	return items
}
