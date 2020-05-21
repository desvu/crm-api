package entity

import "github.com/qilin/crm-api/internal/domain/enum/game_social_link"

type SocialLink struct {
	Type game_social_link.Type
	URL  string
}
