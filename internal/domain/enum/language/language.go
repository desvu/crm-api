package language

type Language string

const (
	English Language = "en"
	Russian Language = "ru"
)

func NewLanguage(v string) Language {
	var source Language

	switch v := Language(v); v {
	case English,
		Russian:
		source = v
	default:
		source = English
	}

	return source
}
