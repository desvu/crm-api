package block

//swagger:enum Title
type Title string

const (
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
	// special case not defined
	if t == "" {
		return true
	}
	switch t {
	case
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
