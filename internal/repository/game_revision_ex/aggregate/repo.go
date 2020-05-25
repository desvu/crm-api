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
	GameMediaRepository             repository.GameMediaRepository
	GameRevisionDeveloperRepository repository.GameRevisionDeveloperRepository
	GameRevisionPublisherRepository repository.GameRevisionPublisherRepository
	GameRevisionFeatureRepository   repository.GameRevisionFeatureRepository
	GameRevisionGenreRepository     repository.GameRevisionGenreRepository
	GameRevisionMediaRepository     repository.GameRevisionMediaRepository
	GameRevisionRatingRepository    repository.GameRevisionRatingRepository
	GameRevisionReviewRepository    repository.GameRevisionReviewRepository
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

func (r GameExRepository) FindLastByGameIDs(ctx context.Context, gameIDs []string) ([]entity.GameRevisionEx, error) {
	games, err := r.GameRevisionRepository.FindLastByGameIDs(ctx, gameIDs)
	if err != nil {
		return nil, err
	}

	if len(games) == 0 {
		return nil, nil
	}

	return r.fetchRows(ctx, games)
}

func (r GameExRepository) FindLastPublishedByGameID(ctx context.Context, gameID string) (*entity.GameRevisionEx, error) {
	game, err := r.GameRevisionRepository.FindLastPublishedByGameID(ctx, gameID)
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
		media         []entity.GameMedia
		localizations []entity.Localization
		ratings       []entity.Rating
		reviews       []entity.Review
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
	g.Go(func() error {
		var err error
		ratings, err = r.GameRevisionRatingRepository.FindByGameRevisionID(ctx, item.ID)
		if err != nil {
			return err
		}
		return nil
	})
	g.Go(func() error {
		relations, err := r.GameRevisionMediaRepository.FindByRevisionID(ctx, item.ID)
		if err != nil {
			return err
		}
		media, err = r.GameMediaRepository.FindByIDs(ctx, entity.NewGameRevisionMediaArray(relations).MediaIDs())
		if err != nil {
			return err
		}
		return nil
	})
	g.Go(func() error {
		var err error
		reviews, err = r.GameRevisionReviewRepository.FindByGameRevisionID(ctx, item.ID)
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
		Media:        media,
		Localization: localizations,
		Rating:       ratings,
		Review:       reviews,
	}, nil
}

func (r GameExRepository) fetchRows(ctx context.Context, items []entity.GameRevision) ([]entity.GameRevisionEx, error) {
	if len(items) == 0 {
		return nil, nil
	}

	var (
		tags               []entity.Tag
		revisionTags       []entity.GameRevisionTag
		developers         []entity.Developer
		revisionDevelopers []entity.GameRevisionDeveloper
		publishers         []entity.Publisher
		revisionPublishers []entity.GameRevisionPublisher
		features           []entity.Feature
		revisionFeatures   []entity.GameRevisionFeature
		genres             []entity.Genre
		revisionGenres     []entity.GameRevisionGenre
		media              []entity.GameMedia
		revisionMedia      []entity.GameRevisionMedia
		localizations      []entity.Localization
		err                error
	)

	g, ctx := errgroup.WithContext(ctx)
	revisionArray := entity.NewGameRevisionArray(items)
	g.Go(func() error {
		revisionTags, err = r.GameRevisionTagRepository.FindByGameRevisionIDs(ctx, revisionArray.IDs())
		if err != nil {
			return err
		}
		tags, err = r.TagRepository.FindByIDs(ctx, entity.NewGameRevisionTagArray(revisionTags).TagIDs())
		if err != nil {
			return err
		}
		return nil
	})
	g.Go(func() error {
		revisionDevelopers, err = r.GameRevisionDeveloperRepository.FindByGameRevisionIDs(ctx, revisionArray.IDs())
		if err != nil {
			return err
		}
		developers, err = r.DeveloperRepository.FindByIDs(ctx, entity.NewGameRevisionDeveloperArray(revisionDevelopers).DeveloperIDs())
		if err != nil {
			return err
		}
		return nil
	})
	g.Go(func() error {
		revisionPublishers, err = r.GameRevisionPublisherRepository.FindByGameRevisionIDs(ctx, revisionArray.IDs())
		if err != nil {
			return err
		}
		publishers, err = r.PublisherRepository.FindByIDs(ctx, entity.NewGameRevisionPublisherArray(revisionPublishers).PublisherIDs())
		if err != nil {
			return err
		}
		return nil
	})
	g.Go(func() error {
		revisionFeatures, err = r.GameRevisionFeatureRepository.FindByGameRevisionIDs(ctx, revisionArray.IDs())
		if err != nil {
			return err
		}
		features, err = r.FeatureRepository.FindByIDs(ctx, entity.NewGameRevisionFeatureArray(revisionFeatures).FeatureIDs())
		if err != nil {
			return err
		}
		return nil
	})
	g.Go(func() error {
		revisionGenres, err = r.GameRevisionGenreRepository.FindByGameRevisionIDs(ctx, revisionArray.IDs())
		if err != nil {
			return err
		}
		genres, err = r.GenreRepository.FindByIDs(ctx, entity.NewGameRevisionGenreArray(revisionGenres).GenreIDs())
		if err != nil {
			return err
		}
		return nil
	})
	g.Go(func() error {
		revisionMedia, err = r.GameRevisionMediaRepository.FindByRevisionIDs(ctx, revisionArray.IDs())
		if err != nil {
			return err
		}
		media, err = r.GameMediaRepository.FindByIDs(ctx, entity.NewGameRevisionMediaArray(revisionMedia).MediaIDs())
		if err != nil {
			return err
		}
		return nil
	})
	g.Go(func() error {
		var err error
		localizations, err = r.LocalizationRepository.FindByGameRevisionIDs(ctx, revisionArray.IDs())
		if err != nil {
			return err
		}
		return nil
	})

	if err := g.Wait(); err != nil {
		return nil, err
	}

	var (
		revisions    = make([]entity.GameRevisionEx, len(items))
		mapRevisions = make(map[uint]*entity.GameRevisionEx)
	)

	for i := range items {
		revisions[i].GameRevision = items[i]
		mapRevisions[items[i].ID] = &revisions[i]
	}

	for i := range tags {
		for j := range revisionTags {
			if tags[i].ID == revisionTags[j].TagID {
				mapRevisions[revisionTags[j].GameRevisionID].Tags = append(
					mapRevisions[revisionTags[j].GameRevisionID].Tags,
					tags[i],
				)
			}
		}
	}

	for i := range developers {
		for j := range revisionDevelopers {
			if developers[i].ID == revisionDevelopers[j].DeveloperID {
				mapRevisions[revisionDevelopers[j].GameRevisionID].Developers = append(
					mapRevisions[revisionDevelopers[j].GameRevisionID].Developers,
					developers[i],
				)
			}
		}
	}

	for i := range publishers {
		for j := range revisionPublishers {
			if publishers[i].ID == revisionPublishers[j].PublisherID {
				mapRevisions[revisionPublishers[j].GameRevisionID].Publishers = append(
					mapRevisions[revisionPublishers[j].GameRevisionID].Publishers,
					publishers[i],
				)
			}
		}
	}

	for i := range features {
		for j := range revisionFeatures {
			if features[i].ID == revisionFeatures[j].FeatureID {
				mapRevisions[revisionFeatures[j].GameRevisionID].Features = append(
					mapRevisions[revisionFeatures[j].GameRevisionID].Features,
					features[i],
				)
			}
		}
	}

	for i := range genres {
		for j := range revisionGenres {
			if genres[i].ID == revisionGenres[j].GenreID {
				mapRevisions[revisionGenres[j].GameRevisionID].Genres = append(
					mapRevisions[revisionGenres[j].GameRevisionID].Genres,
					genres[i],
				)
			}
		}
	}

	for i := range media {
		for j := range revisionMedia {
			if media[i].ID == revisionMedia[j].MediaID {
				mapRevisions[revisionMedia[j].RevisionID].Media = append(
					mapRevisions[revisionMedia[j].RevisionID].Media,
					media[i],
				)
			}
		}
	}

	for i := range localizations {
		mapRevisions[localizations[i].GameRevisionID].Localization = append(
			mapRevisions[localizations[i].GameRevisionID].Localization,
			localizations[i],
		)
	}

	return revisions, nil
}
