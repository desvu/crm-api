package block

type Title string

const (
	Title_None         Title = ""
	Title_SpecialOffer Title = "special-offer"
	Title_GamesOnSale  Title = "games-on-sale"
	Title_NewGames     Title = "new-games"
	Title_MostPopular  Title = "most-popular"
	Title_Trending     Title = "trending"
	Title_YouMayLikeIt Title = "you-may-like-it"
)

func (t Title) String() string {
	return string(t)
}

func (t Title) Valid() bool {
	switch t {
	case
		Title_None,
		Title_SpecialOffer,
		Title_GamesOnSale,
		Title_NewGames,
		Title_MostPopular,
		Title_Trending,
		Title_YouMayLikeIt:
		return true
	}
	return false
}
