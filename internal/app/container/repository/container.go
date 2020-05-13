package repository

import (
	"github.com/qilin/crm-api/internal/repository/developer"
	"github.com/qilin/crm-api/internal/repository/feature"
	"github.com/qilin/crm-api/internal/repository/game"
	"github.com/qilin/crm-api/internal/repository/game_media"
	"github.com/qilin/crm-api/internal/repository/game_revision"
	"github.com/qilin/crm-api/internal/repository/game_revision_developer"
	"github.com/qilin/crm-api/internal/repository/game_revision_ex/aggregate"
	"github.com/qilin/crm-api/internal/repository/game_revision_feature"
	"github.com/qilin/crm-api/internal/repository/game_revision_genre"
	"github.com/qilin/crm-api/internal/repository/game_revision_localization"
	"github.com/qilin/crm-api/internal/repository/game_revision_media"
	"github.com/qilin/crm-api/internal/repository/game_revision_publisher"
	"github.com/qilin/crm-api/internal/repository/game_revision_rating"
	"github.com/qilin/crm-api/internal/repository/game_revision_tag"
	"github.com/qilin/crm-api/internal/repository/game_store_publish"
	"github.com/qilin/crm-api/internal/repository/genre"
	"github.com/qilin/crm-api/internal/repository/publisher"
	"github.com/qilin/crm-api/internal/repository/storefront"
	"github.com/qilin/crm-api/internal/repository/tag"
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
		game_revision_developer.New,
		game_revision_feature.New,
		game_revision_publisher.New,
		game_revision_tag.New,
		game_revision_genre.New,
		aggregate.New,
		game_revision.New,
		game_media.New,
		game_revision_media.New,
		game_store_publish.New,
		storefront.New,
		game_revision_localization.New,
		game_revision_rating.New,
	)
}
