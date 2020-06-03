package service

import (
	"github.com/qilin/crm-api/internal/service/developer"
	"github.com/qilin/crm-api/internal/service/document"
	"github.com/qilin/crm-api/internal/service/feature"
	"github.com/qilin/crm-api/internal/service/game"
	"github.com/qilin/crm-api/internal/service/game_media"
	"github.com/qilin/crm-api/internal/service/game_revision"
	"github.com/qilin/crm-api/internal/service/game_store_publish"
	"github.com/qilin/crm-api/internal/service/genre"
	"github.com/qilin/crm-api/internal/service/localization"
	"github.com/qilin/crm-api/internal/service/publisher"
	"github.com/qilin/crm-api/internal/service/rating"
	"github.com/qilin/crm-api/internal/service/review"
	"github.com/qilin/crm-api/internal/service/storefront"
	"github.com/qilin/crm-api/internal/service/tag"
	"github.com/qilin/crm-api/internal/service/user"
	"go.uber.org/fx"
)

func New() fx.Option {
	return fx.Provide(
		developer.New,
		feature.New,
		game.New,
		genre.New,
		publisher.New,
		tag.New,
		game_revision.New,
		game_media.New,
		game_store_publish.New,
		storefront.New,
		localization.New,
		rating.New,
		review.New,
		document.New,
		user.New,
	)
}
