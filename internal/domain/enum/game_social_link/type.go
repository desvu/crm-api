package game_social_link

type Type uint8

const (
	TypeUndefined Type = 0
	TypeTwitch    Type = 1
	TypeFacebook  Type = 2
	TypeYoutube   Type = 3
	TypeTwitter   Type = 4
	TypeDiscord   Type = 5
	TypeReddit    Type = 6
	TypeVk        Type = 7
)

func NewType(v uint8) Type {
	var source Type

	switch v := Type(v); v {
	case TypeUndefined,
		TypeTwitch,
		TypeFacebook,
		TypeYoutube,
		TypeTwitter,
		TypeDiscord,
		TypeReddit,
		TypeVk:
		source = v
	}

	return source
}

func NewTypeByString(v string) Type {
	switch v {
	case "twitch":
		return TypeTwitch
	case "facebook":
		return TypeFacebook
	case "youtube":
		return TypeYoutube
	case "twitter":
		return TypeTwitter
	case "discord":
		return TypeDiscord
	case "reddit":
		return TypeReddit
	case "vk":
		return TypeVk
	default:
		return TypeUndefined
	}
}

func (t Type) Value() uint8 {
	return uint8(t)
}

func (t Type) String() string {
	switch t {
	case TypeTwitch:
		return "twitch"
	case TypeFacebook:
		return "facebook"
	case TypeYoutube:
		return "youtube"
	case TypeTwitter:
		return "twitter"
	case TypeDiscord:
		return "discord"
	case TypeReddit:
		return "reddit"
	case TypeVk:
		return "vk"
	default:
		return "undefined"
	}
}
