package service

import (
	"github.com/qilin/crm-api/internal/app/repository"
	"github.com/qilin/crm-api/internal/domain/service"
	"github.com/qilin/crm-api/internal/service/developer"
	"github.com/qilin/crm-api/internal/service/feature"
	"github.com/qilin/crm-api/internal/service/game"
	"github.com/qilin/crm-api/internal/service/genre"
	"github.com/qilin/crm-api/internal/service/publisher"
	"github.com/qilin/crm-api/internal/service/tag"
)

type Services struct {
	GameService      service.GameService
	TagService       service.TagService
	DeveloperService service.DeveloperService
	PublisherService service.PublisherService
	FeatureService   service.FeatureService
	GenreService     service.GenreService
}

func New(r *repository.Repositories) *Services {
	s := new(Services)

	s.GameService = game.New(
		game.ServiceParams{
			GameRepository: r.GameRepository,
		},
	)

	s.TagService = tag.New(
		tag.ServiceParams{
			GameService:       s.GameService,
			TagRepository:     r.TagRepository,
			GameTagRepository: r.GameTagRepository,
		},
	)

	s.DeveloperService = developer.New(
		developer.ServiceParams{
			GameService:             s.GameService,
			DeveloperRepository:     r.DeveloperRepository,
			GameDeveloperRepository: r.GameDeveloperRepository,
		},
	)

	s.PublisherService = publisher.New(
		publisher.ServiceParams{
			GameService:             s.GameService,
			PublisherRepository:     r.PublisherRepository,
			GamePublisherRepository: r.GamePublisherRepository,
		},
	)

	s.FeatureService = feature.New(
		feature.ServiceParams{
			GameService:           s.GameService,
			FeatureRepository:     r.FeatureRepository,
			GameFeatureRepository: r.GameFeatureRepository,
		},
	)

	s.GenreService = genre.New(
		genre.ServiceParams{
			GameService:         s.GameService,
			GenreRepository:     r.GenreRepository,
			GameGenreRepository: r.GameGenreRepository,
		},
	)

	return s
}
