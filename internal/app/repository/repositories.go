package repository

import (
	"github.com/qilin/crm-api/internal/domain/repository"
	"github.com/qilin/crm-api/internal/env"
	"github.com/qilin/crm-api/internal/repository/developer"
	"github.com/qilin/crm-api/internal/repository/feature"
	"github.com/qilin/crm-api/internal/repository/game"
	"github.com/qilin/crm-api/internal/repository/game_developer"
	"github.com/qilin/crm-api/internal/repository/game_feature"
	"github.com/qilin/crm-api/internal/repository/game_genre"
	"github.com/qilin/crm-api/internal/repository/game_publisher"
	"github.com/qilin/crm-api/internal/repository/game_tag"
	"github.com/qilin/crm-api/internal/repository/genre"
	"github.com/qilin/crm-api/internal/repository/publisher"
	"github.com/qilin/crm-api/internal/repository/tag"
)

type Repositories struct {
	GameRepository          repository.GameRepository
	TagRepository           repository.TagRepository
	GameTagRepository       repository.GameTagRepository
	DeveloperRepository     repository.DeveloperRepository
	GameDeveloperRepository repository.GameDeveloperRepository
	PublisherRepository     repository.PublisherRepository
	GamePublisherRepository repository.GamePublisherRepository
	FeatureRepository       repository.FeatureRepository
	GameFeatureRepository   repository.GameFeatureRepository
	GenreRepository         repository.GenreRepository
	GameGenreRepository     repository.GameGenreRepository
}

func New(e *env.Store) *Repositories {
	return &Repositories{
		GameRepository:          game.New(e),
		TagRepository:           tag.New(e),
		GameTagRepository:       game_tag.New(e),
		DeveloperRepository:     developer.New(e),
		GameDeveloperRepository: game_developer.New(e),
		PublisherRepository:     publisher.New(e),
		GamePublisherRepository: game_publisher.New(e),
		FeatureRepository:       feature.New(e),
		GameFeatureRepository:   game_feature.New(e),
		GenreRepository:         genre.New(e),
		GameGenreRepository:     game_genre.New(e),
	}
}
