package game_social_link

type Type uint8

const (
	TypeUndefined Type = 1
	TypeTwitch    Type = 2
	TypeFacebook  Type = 3
	TypeYoutube   Type = 4
	TypeTwitter   Type = 5
	TypeDiscord   Type = 6
	TypeReddit    Type = 7
	TypeVk        Type = 8
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
