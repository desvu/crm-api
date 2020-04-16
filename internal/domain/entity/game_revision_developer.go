package entity

type GameRevisionDeveloper struct {
	ID             uint
	GameRevisionID uint
	DeveloperID    uint
}

type GameRevisionDeveloperArray struct {
	ids          []uint
	developerIDs []uint
	items        []GameRevisionDeveloper
}

func NewGameRevisionDeveloperArray(items []GameRevisionDeveloper) *GameRevisionDeveloperArray {
	a := &GameRevisionDeveloperArray{
		items: items,
	}
	a.refresh()
	return a
}

func (a *GameRevisionDeveloperArray) refresh() {
	ids := make([]uint, len(a.items))
	developerIDs := make([]uint, len(a.items))

	for i := range a.items {
		ids[i] = a.items[i].ID
		developerIDs[i] = a.items[i].DeveloperID
	}

	a.ids = ids
	a.developerIDs = developerIDs
}

func (a *GameRevisionDeveloperArray) Size() int {
	return len(a.items)
}

func (a *GameRevisionDeveloperArray) IDs() []uint {
	items := make([]uint, len(a.items))
	copy(items, a.ids)
	return items
}

func (a *GameRevisionDeveloperArray) DeveloperIDs() []uint {
	items := make([]uint, len(a.items))
	copy(items, a.developerIDs)
	return items
}
