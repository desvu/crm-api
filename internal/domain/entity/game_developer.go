package entity

type GameDeveloper struct {
	ID          uint
	GameID      uint
	DeveloperID uint
}

type GameDeveloperArray struct {
	ids          []uint
	developerIDs []uint
	items        []GameDeveloper
}

func NewGameDeveloperArray(items []GameDeveloper) *GameDeveloperArray {
	a := &GameDeveloperArray{
		items: items,
	}
	a.refresh()
	return a
}

func (a *GameDeveloperArray) refresh() {
	ids := make([]uint, len(a.items))
	developerIDs := make([]uint, len(a.items))

	for i := range a.items {
		ids[i] = a.items[i].ID
		developerIDs[i] = a.items[i].DeveloperID
	}

	a.ids = ids
	a.developerIDs = developerIDs
}

func (a *GameDeveloperArray) Size() int {
	return len(a.items)
}

func (a *GameDeveloperArray) IDs() []uint {
	items := make([]uint, len(a.items))
	copy(items, a.ids)
	return items
}

func (a *GameDeveloperArray) DeveloperIDs() []uint {
	items := make([]uint, len(a.items))
	copy(items, a.developerIDs)
	return items
}
