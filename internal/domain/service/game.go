package service

import (
	"context"
	"net/url"
	"regexp"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/qilin/crm-api/internal/domain/entity"
	"github.com/qilin/crm-api/internal/domain/enum/enum"
	"github.com/qilin/crm-api/internal/domain/enum/game"
	"github.com/qilin/crm-api/internal/domain/enum/game_social_link"
	"github.com/qilin/crm-api/internal/domain/errors"
)

//go:generate mockgen -destination=../mocks/game_service.go -package=mocks github.com/qilin/crm-api/internal/domain/service GameService
type GameService interface {
	Create(ctx context.Context, data *CreateGameData) (*entity.GameEx, error)
	Update(ctx context.Context, data *UpdateGameData) (*entity.GameEx, error)
	Upsert(ctx context.Context, data *UpsertGameData) (*entity.GameEx, error)
	Delete(ctx context.Context, id string) error
	Publish(ctx context.Context, id string) error

	GetByID(ctx context.Context, id string) (*entity.Game, error)
	GetExByID(ctx context.Context, id string) (*entity.GameEx, error)

	// GetExLastPublishedByID returns last published game by id
	GetExLastPublishedByID(ctx context.Context, id string) (*entity.GameEx, error)

	// GetExBySlug returns last published game by slug
	GetExBySlug(ctx context.Context, slug string) (*entity.GameEx, error)

	// FindByTitle returns games found by title substring
	GetByTitleSubstring(ctx context.Context, data GetByTitleSubstringData) ([]entity.GameEx, error)

	// GetExByFilter returns games founded by filter
	GetExByFilter(ctx context.Context, data *GetByFilterGameData) ([]entity.GameEx, error)

	// GetCountByFilter returns count games founded by filter
	GetCountByFilter(ctx context.Context, data *GetByFilterGameData) (int, error)
}

type CommonGameData struct {
	Summary     *string
	Description *string
	License     *string
	Trailer     *string // `validate:"trailer"`
	PlayTime    *uint
	Tags        *[]uint
	Developers  *[]uint
	Publishers  *[]uint
	Features    *[]uint
	Genres      *[]uint
	Media       *[]uint

	SocialLinks        *[]SocialLink
	SystemRequirements *[]SystemRequirements
	Platforms          *game.PlatformArray
	ReleaseDate        *time.Time
	Localizations      *[]LocalizationData
	Ratings            *[]RatingData
	Reviews            *[]ReviewData
}

type UpsertGameData struct {
	ID    *string
	Title *string
	Slug  *string
	Type  *game.Type

	CommonGameData
}

type CreateGameData struct {
	Title string    `validate:"required"`
	Slug  string    `validate:"required"`
	Type  game.Type `validate:"required"`

	CommonGameData
}

func (d CreateGameData) Validate() error {
	validate := validator.New()

	//err := validate.RegisterValidation("trailer", validateTrailer)
	//if err != nil {
	//	return err
	//}

	err := validate.Struct(d)
	if err != nil {
		return err
	}

	if d.Localizations != nil {
		for _, l := range *d.Localizations {
			if err := l.Validate(); err != nil {
				return err
			}
		}
	}

	if d.SocialLinks != nil {
		for _, l := range *d.SocialLinks {
			if err := l.Validate(); err != nil {
				return err
			}
		}
	}

	if d.Ratings != nil {
		for _, l := range *d.Ratings {
			if err := l.Validate(); err != nil {
				return err
			}
		}
	}

	if d.Reviews != nil {
		if len(*d.Reviews) > 3 {
			return errors.ReviewMax3Available
		}
		for _, r := range *d.Reviews {
			if err := r.Validate(); err != nil {
				return err
			}
		}
	}

	return nil
}

type UpdateGameData struct {
	ID    string
	Title *string
	Slug  *string
	Type  *game.Type

	CommonGameData
}

func (d UpdateGameData) Validate() error {
	if d.Localizations != nil {
		for _, l := range *d.Localizations {
			if err := l.Validate(); err != nil {
				return err
			}
		}
	}
	if d.Reviews != nil {
		if len(*d.Reviews) > 3 {
			return errors.ReviewMax3Available
		}
		for _, r := range *d.Reviews {
			if err := r.Validate(); err != nil {
				return err
			}
		}
	}
	return nil
}

type SocialLink struct {
	Type game_social_link.Type
	URL  string
}

func (d SocialLink) Validate() error {
	u, err := url.Parse(d.URL)
	if err != nil {
		return errors.SocialLinkIncorrectURL
	}

	if u.Hostname() == "" {
		return errors.SocialLinkIncorrectURL
	}

	return nil
}

type SystemRequirements struct {
	Platform    game.Platform
	Minimal     *RequirementsSet
	Recommended *RequirementsSet
}

type RequirementsSet struct {
	OS        string
	CPU       string
	GPU       string
	DiskSpace uint
	RAM       uint
}

func validateTrailer(fl validator.FieldLevel) bool {
	match, _ := regexp.MatchString(
		`^((?:https?:)?\/\/)?((?:www|m)\.)?((?:youtube\.com|youtu.be))(\/(?:[\w\-]+\?v=|embed\/|v\/)?)([\w\-]+)(\S+)?$`,
		fl.Field().String(),
	)

	return match
}

type GetByFilterGameData struct {
	OnlyPublished bool
	GenreIDs      []uint
	FeatureIDs    []uint
	Languages     []string
	Platforms     []game.Platform
	OrderType     enum.SortOrderType
	OrderBy       enum.SortOrderColumn
	Limit         int
	Offset        int
}

type GetByTitleSubstringData struct {
	Title  string
	Limit  int
	Offset int
}
