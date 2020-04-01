package game

type PlatformArray []Platform

func NewPlatformArray(v ...uint8) PlatformArray {
	source := make(PlatformArray, 0, len(v))

	for i := range v {
		source.Add(NewPlatform(v[i]))
	}

	return source
}

func (p *PlatformArray) Add(v Platform) {
	if p == nil {
		return
	}

	for _, platform := range *p {
		if platform == v {
			return
		}
	}
	*p = append(*p, v)
}

func (p PlatformArray) Values() []uint8 {
	source := make([]uint8, len(p))
	for i, platform := range p {
		source[i] = platform.Value()
	}

	return source
}
