package entity

type GameEx struct {
	Game
	Tags         []Tag
	Developers   []Developer
	Publishers   []Publisher
	Features     []Feature
	Genres       []Genre
	Languages    []Language
	Ratings      []Rating
	Requirements []SystemRequirements
}
