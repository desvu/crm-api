package game

type PlatformArray struct {
	len   int
	items []Platform
}

func NewPlatformArray(v []uint8) PlatformArray {
	source := PlatformArray{}

	for i := range v {
		source.Add(NewPlatform(v[i]))
	}

	return source
}

func (p *PlatformArray) Add(v Platform) {
	if p == nil {
		return
	}

	for i := range p.items {
		if (*p).items[i] == v {
			return
		}
	}

	p.len++
	p.items = append(p.items, v)
}

func (p *PlatformArray) Values() []uint8 {
	source := make([]uint8, p.len)
	for i := range p.items {
		source[i] = p.items[i].Value()
	}

	return source
}
