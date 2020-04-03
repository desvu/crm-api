package aggregate

import (
	"context"

	"github.com/qilin/crm-api/internal/domain/entity"
	"github.com/qilin/crm-api/internal/domain/repository"
	"golang.org/x/sync/errgroup"
)

type RepositoryParams struct {
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

type GameExRepository struct {
	RepositoryParams
}

func New(params RepositoryParams) repository.GameExRepository {
	return &GameExRepository{
		params,
	}
}

func (r GameExRepository) FindByID(ctx context.Context, id uint) (*entity.GameEx, error) {
	game, err := r.GameRepository.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if game == nil {
		return nil, nil
	}

	return r.fetchRow(ctx, game)
}

func (r GameExRepository) fetchRow(ctx context.Context, item *entity.Game) (*entity.GameEx, error) {
	if item == nil {
		return nil, nil
	}

	var (
		tags       []entity.Tag
		developers []entity.Developer
		publishers []entity.Publisher
		features   []entity.Feature
		genres     []entity.Genre
	)

	g, ctx := errgroup.WithContext(ctx)
	g.Go(func() error {
		relations, err := r.GameTagRepository.FindByGameID(ctx, item.ID)
		if err != nil {
			return err
		}
		tags, err = r.TagRepository.FindByIDs(ctx, entity.NewGameTagArray(relations).TagIDs())
		return nil
	})
	g.Go(func() error {
		relations, err := r.GameDeveloperRepository.FindByGameID(ctx, item.ID)
		if err != nil {
			return err
		}
		developers, err = r.DeveloperRepository.FindByIDs(ctx, entity.NewGameDeveloperArray(relations).DeveloperIDs())
		return nil
	})
	g.Go(func() error {
		relations, err := r.GamePublisherRepository.FindByGameID(ctx, item.ID)
		if err != nil {
			return err
		}
		publishers, err = r.PublisherRepository.FindByIDs(ctx, entity.NewGamePublisherArray(relations).PublisherIDs())
		return nil
	})
	g.Go(func() error {
		relations, err := r.GameFeatureRepository.FindByGameID(ctx, item.ID)
		if err != nil {
			return err
		}
		features, err = r.FeatureRepository.FindByIDs(ctx, entity.NewGameFeatureArray(relations).FeatureIDs())
		return nil
	})
	g.Go(func() error {
		relations, err := r.GameGenreRepository.FindByGameID(ctx, item.ID)
		if err != nil {
			return err
		}
		genres, err = r.GenreRepository.FindByIDs(ctx, entity.NewGameGenreArray(relations).GenreIDs())
		return nil
	})

	if err := g.Wait(); err != nil {
		return nil, err
	}

	return &entity.GameEx{
		Game:       *item,
		Tags:       tags,
		Developers: developers,
		Publishers: publishers,
		Features:   features,
		Genres:     genres,
	}, nil
}
