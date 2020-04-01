package entity

type GameExt struct {
	Game
	Developers   []Developer
	Publishers   []Publisher
	Genres       []Genre
	Tags         []Tag
	Languages    []Language
	Features     []Feature
	Ratings      []Rating
	Requirements []SystemRequirements
}
