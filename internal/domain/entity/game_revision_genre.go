package entity

type GameRevisionGenre struct {
	ID             uint
	GameRevisionID uint
	GenreID        uint
}

type GameRevisionGenreArray struct {
	ids      []uint
	genreIDs []uint
	items    []GameRevisionGenre
}

func NewGameRevisionGenreArray(items []GameRevisionGenre) *GameRevisionGenreArray {
	a := &GameRevisionGenreArray{
		items: items,
	}
	a.refresh()
	return a
}

func (a *GameRevisionGenreArray) refresh() {
	ids := make([]uint, len(a.items))
	genreIDs := make([]uint, len(a.items))

	for i := range a.items {
		ids[i] = a.items[i].ID
		genreIDs[i] = a.items[i].GenreID
	}

	a.ids = ids
	a.genreIDs = genreIDs
}

func (a *GameRevisionGenreArray) Size() int {
	return len(a.items)
}

func (a *GameRevisionGenreArray) IDs() []uint {
	items := make([]uint, len(a.items))
	copy(items, a.ids)
	return items
}

func (a *GameRevisionGenreArray) GenreIDs() []uint {
	items := make([]uint, len(a.items))
	copy(items, a.genreIDs)
	return items
}
