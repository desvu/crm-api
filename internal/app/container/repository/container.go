package repository

import (
	"github.com/qilin/crm-api/internal/repository/developer"
	"github.com/qilin/crm-api/internal/repository/feature"
	"github.com/qilin/crm-api/internal/repository/game"
	"github.com/qilin/crm-api/internal/repository/game_developer"
	"github.com/qilin/crm-api/internal/repository/game_ex/aggregate"
	"github.com/qilin/crm-api/internal/repository/game_feature"
	"github.com/qilin/crm-api/internal/repository/game_genre"
	"github.com/qilin/crm-api/internal/repository/game_publisher"
	"github.com/qilin/crm-api/internal/repository/game_tag"
	"github.com/qilin/crm-api/internal/repository/genre"
	"github.com/qilin/crm-api/internal/repository/publisher"
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
		game_developer.New,
		game_feature.New,
		game_publisher.New,
		game_tag.New,
		game_genre.New,
		aggregate.New,
	)
}
