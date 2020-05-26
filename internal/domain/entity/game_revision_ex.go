package entity

type GameRevisionEx struct {
	GameRevision
	Tags         []Tag
	Developers   []Developer
	Publishers   []Publisher
	Features     []Feature
	Genres       []Genre
	Media        []GameMedia
	Localization []Localization
	Rating       []Rating
	Review       []Review
}

type GameRevisionExArray []GameRevisionEx

func NewGameRevisionExArray(items []GameRevisionEx) GameRevisionExArray {
	return GameRevisionExArray(items)
}

func (a GameRevisionExArray) IDs() []uint {
	ids := make([]uint, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}

	return ids
}

func (a GameRevisionExArray) GameIDs() []string {
	ids := make([]string, len(a))
	for i := range a {
		ids[i] = a[i].GameID
	}

	return ids
}
