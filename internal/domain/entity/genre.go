package entity

type Genre struct {
	ID   uint
	Name string
}

type GenreArray struct {
	ids   []uint
	items []Genre
}

func NewGenreArray(items []Genre) *GenreArray {
	a := &GenreArray{
		items: items,
	}
	a.refresh()
	return a
}

func (a *GenreArray) refresh() {
	ids := make([]uint, len(a.items))

	for i := range a.items {
		ids[i] = a.items[i].ID
	}

	a.ids = ids
}

func (a *GenreArray) Size() int {
	return len(a.items)
}

func (a *GenreArray) IDs() []uint {
	items := make([]uint, len(a.items))
	copy(items, a.ids)
	return items
}
