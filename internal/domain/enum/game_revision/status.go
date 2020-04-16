package game_revision

type Status uint8

const (
	StatusDraft     Status = 1
	StatusSkipped   Status = 2
	StatusPublished Status = 3
)

func NewStatus(v uint8) Status {
	var source Status

	switch v := Status(v); v {
	case StatusDraft,
		StatusSkipped,
		StatusPublished:
		source = v
	}

	return source
}

func (t Status) Value() uint8 {
	return uint8(t)
}

func (t Status) String() string {
	var str string

	switch t {
	case StatusDraft:
		str = "draft"
	case StatusSkipped:
		str = "skipped"
	case StatusPublished:
		str = "published"
	}

	return str
}
