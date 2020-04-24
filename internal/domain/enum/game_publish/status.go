package game_publish

type Status uint8

const (
	StatusUndefined Status = 0
	StatusInProcess Status = 1
	StatusPublished Status = 2
	StatusFailure   Status = 3
)

func NewStatus(v uint8) Status {
	var source Status

	switch v := Status(v); v {
	case StatusUndefined,
		StatusInProcess,
		StatusPublished,
		StatusFailure:
		source = v
	}

	return source
}

func (t Status) Value() uint8 {
	return uint8(t)
}

func (t Status) String() string {
	switch t {
	case StatusInProcess:
		return "in_process"
	case StatusPublished:
		return "published"
	case StatusFailure:
		return "failure"
	default:
		return "undefined"
	}
}
