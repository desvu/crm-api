package block

//swagger:enum Title
type Title string

const (
	TitleUndefined    Title = "undefined"
	TitleSpecialOffer Title = "special-offer"
	TitleGamesOnSale  Title = "games-on-sale"
	TitleNewGames     Title = "new-games"
	TitleMostPopular  Title = "most-popular"
	TitleTrending     Title = "trending"
	TitleYouMayLikeIt Title = "you-may-like-it"
)

func (t Title) String() string {
	return string(t)
}

func (t Title) Valid() bool {
	switch t {
	case
		TitleUndefined,
		TitleSpecialOffer,
		TitleGamesOnSale,
		TitleNewGames,
		TitleMostPopular,
		TitleTrending,
		TitleYouMayLikeIt:
		return true
	}
	return false
}
