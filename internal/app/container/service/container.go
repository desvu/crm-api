package service

import (
	"github.com/qilin/crm-api/internal/service/developer"
	"github.com/qilin/crm-api/internal/service/feature"
	"github.com/qilin/crm-api/internal/service/game"
	"github.com/qilin/crm-api/internal/service/genre"
	"github.com/qilin/crm-api/internal/service/publisher"
	"github.com/qilin/crm-api/internal/service/tag"
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
	)
}
