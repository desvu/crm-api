package errors

import "github.com/qilin/crm-api/pkg/errors"

func NewValidation(err error) error {
	return errors.NewService(errors.ErrValidation, err.Error())
}

var (
	// Developer service
	DeveloperNotFound   = errors.NewService(errors.ErrNotFound, "developer not found")
	InvalidDeveloperIDs = errors.NewService(errors.ErrValidation, "invalid developer ids")

	// Feature service
	FeatureNotFound   = errors.NewService(errors.ErrNotFound, "feature not found")
	InvalidFeatureIDs = errors.NewService(errors.ErrValidation, "invalid feature ids")

	// Game service
	GameSlugAlreadyExist = errors.NewService(errors.ErrAlreadyExist, "game slug already exist")
	GameNotFound         = errors.NewService(errors.ErrNotFound, "game not found")

	// GameRevision service
	GameRevisionNotFound = errors.NewService(errors.ErrNotFound, "game revision not found")

	// GameStorePublish service
	GameStorePublishNotFound = errors.NewService(errors.ErrNotFound, "game store publish not found")

	// Genre service
	GenreNotFound   = errors.NewService(errors.ErrNotFound, "genre not found")
	InvalidGenreIDs = errors.NewService(errors.ErrValidation, "invalid genre ids")

	// Publisher service
	PublisherNotFound   = errors.NewService(errors.ErrNotFound, "publisher not found")
	InvalidPublisherIDs = errors.NewService(errors.ErrValidation, "invalid publisher ids")

	// Tag service
	TagNotFound   = errors.NewService(errors.ErrNotFound, "tag not found")
	InvalidTagIDs = errors.NewService(errors.ErrValidation, "invalid tag ids")

	// GameMedia service
	InvalidMediaIDs = errors.NewService(errors.ErrValidation, "invalid media ids")

	// StoreFront service
	StoreFrontNotFound     = errors.NewService(errors.ErrNotFound, "storefront not found")
	StoreFrontIsActive     = errors.NewService(errors.ErrValidation, "storefront is active")
	UnknownBlockType       = errors.NewService(errors.ErrValidation, "unknown block type")
	InvalidBlockTitle      = errors.NewService(errors.ErrValidation, "invalid block title")
	InvalidBlockGamesCount = errors.NewService(errors.ErrValidation, "invalid block games count")

    // Localization service
    LocalizationNotFound            = errors.NewService(errors.ErrNotFound, "localization not found")
    InvalidLocalizationLanguageCode = errors.NewService(errors.ErrValidation, "invalid localization language code")
)

func NewInternal(err error) errors.Error {
	return errors.NewInternal(err)
}
