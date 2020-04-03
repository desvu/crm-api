package entity

type GameGenre struct {
	ID      uint
	GameID  uint
	GenreID uint
}

type GameGenreArray struct {
	ids      []uint
	genreIDs []uint
	items    []GameGenre
}

func NewGameGenreArray(items []GameGenre) *GameGenreArray {
	a := &GameGenreArray{
		items: items,
	}
	a.refresh()
	return a
}

func (a *GameGenreArray) refresh() {
	ids := make([]uint, len(a.items))
	genreIDs := make([]uint, len(a.items))

	for i := range a.items {
		ids[i] = a.items[i].ID
		genreIDs[i] = a.items[i].GenreID
	}

	a.ids = ids
	a.genreIDs = genreIDs
}

func (a *GameGenreArray) Size() int {
	return len(a.items)
}

func (a *GameGenreArray) IDs() []uint {
	items := make([]uint, len(a.items))
	copy(items, a.ids)
	return items
}

func (a *GameGenreArray) GenreIDs() []uint {
	items := make([]uint, len(a.items))
	copy(items, a.genreIDs)
	return items
}
