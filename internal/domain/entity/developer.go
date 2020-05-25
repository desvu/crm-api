package entity

type Developer struct {
	ID   uint
	Name string
}

type DeveloperArray struct {
	ids   []uint
	items []Developer
}

func NewDeveloperArray(items []Developer) *DeveloperArray {
	a := &DeveloperArray{
		items: items,
	}
	a.refresh()
	return a
}

func (a *DeveloperArray) refresh() {
	ids := make([]uint, len(a.items))

	for i := range a.items {
		ids[i] = a.items[i].ID
	}

	a.ids = ids
}

func (a *DeveloperArray) Size() int {
	return len(a.items)
}

func (a *DeveloperArray) IDs() []uint {
	items := make([]uint, len(a.items))
	copy(items, a.ids)
	return items
}
