package aggregate

import (
	"context"

	"github.com/qilin/crm-api/internal/domain/entity"
	"github.com/qilin/crm-api/internal/domain/repository"
	"go.uber.org/fx"
	"golang.org/x/sync/errgroup"
)

type RepositoryParams struct {
	fx.In

	GameRevisionRepository          repository.GameRevisionRepository
	TagRepository                   repository.TagRepository
	GameRevisionTagRepository       repository.GameRevisionTagRepository
	DeveloperRepository             repository.DeveloperRepository
	PublisherRepository             repository.PublisherRepository
	FeatureRepository               repository.FeatureRepository
	GenreRepository                 repository.GenreRepository
	GameRevisionDeveloperRepository repository.GameRevisionDeveloperRepository
	GameRevisionPublisherRepository repository.GameRevisionPublisherRepository
	GameRevisionFeatureRepository   repository.GameRevisionFeatureRepository
	GameRevisionGenreRepository     repository.GameRevisionGenreRepository
	LocalizationRepository          repository.GameRevisionLocalizationRepository
}

type GameExRepository struct {
	RepositoryParams
}

func New(params RepositoryParams) repository.GameRevisionExRepository {
	return &GameExRepository{
		params,
	}
}

func (r GameExRepository) FindByID(ctx context.Context, id uint) (*entity.GameRevisionEx, error) {
	game, err := r.GameRevisionRepository.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if game == nil {
		return nil, nil
	}

	return r.fetchRow(ctx, game)
}

func (r GameExRepository) FindByIDAndGameID(ctx context.Context, id uint, gameID string) (*entity.GameRevisionEx, error) {
	game, err := r.GameRevisionRepository.FindByIDAndGameID(ctx, id, gameID)
	if err != nil {
		return nil, err
	}

	if game == nil {
		return nil, nil
	}

	return r.fetchRow(ctx, game)
}

func (r GameExRepository) fetchRow(ctx context.Context, item *entity.GameRevision) (*entity.GameRevisionEx, error) {
	if item == nil {
		return nil, nil
	}

	var (
		tags          []entity.Tag
		developers    []entity.Developer
		publishers    []entity.Publisher
		features      []entity.Feature
		genres        []entity.Genre
		localizations []entity.Localization
	)

	g, ctx := errgroup.WithContext(ctx)
	g.Go(func() error {
		relations, err := r.GameRevisionTagRepository.FindByGameRevisionID(ctx, item.ID)
		if err != nil {
			return err
		}
		tags, err = r.TagRepository.FindByIDs(ctx, entity.NewGameRevisionTagArray(relations).TagIDs())
		if err != nil {
			return err
		}
		return nil
	})
	g.Go(func() error {
		relations, err := r.GameRevisionDeveloperRepository.FindByGameRevisionID(ctx, item.ID)
		if err != nil {
			return err
		}
		developers, err = r.DeveloperRepository.FindByIDs(ctx, entity.NewGameRevisionDeveloperArray(relations).DeveloperIDs())
		if err != nil {
			return err
		}
		return nil
	})
	g.Go(func() error {
		relations, err := r.GameRevisionPublisherRepository.FindByGameRevisionID(ctx, item.ID)
		if err != nil {
			return err
		}
		publishers, err = r.PublisherRepository.FindByIDs(ctx, entity.NewGameRevisionPublisherArray(relations).PublisherIDs())
		if err != nil {
			return err
		}
		return nil
	})
	g.Go(func() error {
		relations, err := r.GameRevisionFeatureRepository.FindByGameRevisionID(ctx, item.ID)
		if err != nil {
			return err
		}
		features, err = r.FeatureRepository.FindByIDs(ctx, entity.NewGameRevisionFeatureArray(relations).FeatureIDs())
		if err != nil {
			return err
		}
		return nil
	})
	g.Go(func() error {
		relations, err := r.GameRevisionGenreRepository.FindByGameRevisionID(ctx, item.ID)
		if err != nil {
			return err
		}
		genres, err = r.GenreRepository.FindByIDs(ctx, entity.NewGameRevisionGenreArray(relations).GenreIDs())
		if err != nil {
			return err
		}
		return nil
	})
	g.Go(func() error {
		var err error
		localizations, err = r.LocalizationRepository.FindByGameRevisionID(ctx, item.ID)
		if err != nil {
			return err
		}
		return nil
	})

	if err := g.Wait(); err != nil {
		return nil, err
	}

	return &entity.GameRevisionEx{
		GameRevision: *item,
		Tags:         tags,
		Developers:   developers,
		Publishers:   publishers,
		Features:     features,
		Genres:       genres,
		Localization: localizations,
	}, nil
}
