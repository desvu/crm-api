package errors

import "github.com/qilin/crm-api/pkg/errors"

var (
	// Common
	InvalidGameID = errors.NewService(errors.ErrValidation, "invalid game id", "invalid_game_id")

	// Developer service
	DeveloperNotFound   = errors.NewService(errors.ErrNotFound, "developer not found", "developer_not_found")
	InvalidDeveloperIDs = errors.NewService(errors.ErrValidation, "invalid developer ids", "invalid_developer_ids")

	// Feature service
	FeatureNotFound   = errors.NewService(errors.ErrNotFound, "feature not found", "feature_not_found")
	InvalidFeatureIDs = errors.NewService(errors.ErrValidation, "invalid feature ids", "invalid_feature_ids")

	// Game service
	GameSlugAlreadyExist = errors.NewService(errors.ErrAlreadyExist, "game slug already exist", "game_slug_already_exist")
	GameNotFound         = errors.NewService(errors.ErrNotFound, "game not found", "game_not_found")

	// GameRevision service
	GameRevisionNotFound                 = errors.NewService(errors.ErrNotFound, "game revision not found", "game_revision_not_found")
	GameRevisionUniqueSystemRequirements = errors.NewService(errors.ErrValidation, "systemRequirements platform param must be unique", "game_revision_unique_system_requirements")
	GameRevisionDraftAlreadyExist        = errors.NewService(errors.ErrAlreadyExist, "game revision draft already exist", "game_revision_draft_already_exist")

	// GameStorePublish service
	GameStorePublishNotFound = errors.NewService(errors.ErrNotFound, "game store publish not found", "game_store_publish_not_found")

	// Genre service
	GenreNotFound   = errors.NewService(errors.ErrNotFound, "genre not found", "genre_not_found")
	InvalidGenreIDs = errors.NewService(errors.ErrValidation, "invalid genre ids", "invalid_genre_ids")

	// Publisher service
	PublisherNotFound   = errors.NewService(errors.ErrNotFound, "publisher not found", "publisher_not_found")
	InvalidPublisherIDs = errors.NewService(errors.ErrValidation, "invalid publisher ids", "invalid_publisher_ids")

	// Tag service
	TagNotFound   = errors.NewService(errors.ErrNotFound, "tag not found", "tag_not_found")
	InvalidTagIDs = errors.NewService(errors.ErrValidation, "invalid tag ids", "invalid_tag_ids")

	// GameMedia service
	MediaNotFound           = errors.NewService(errors.ErrNotFound, "media not found", "media_not_found")
	InvalidMediaIDs         = errors.NewService(errors.ErrValidation, "invalid media ids", "invalid_media_ids")
	InvalidMediaMIMEType    = errors.NewService(errors.ErrValidation, "invalid media mime type", "invalid_media_mime_type")
	InvalidMediaResolution  = errors.NewService(errors.ErrValidation, "invalid media resolution", "invalid_media_resolution")
	InvalidMediaAspectRatio = errors.NewService(errors.ErrValidation, "invalid media aspect ratio", "invalid_media_aspect_ratio")

	// StoreFront service
	StoreFrontNotFound     = errors.NewService(errors.ErrNotFound, "storefront not found", "storefront_not_found")
	StoreFrontIsActive     = errors.NewService(errors.ErrValidation, "storefront is active", "storefront_is_active")
	UnknownBlockType       = errors.NewService(errors.ErrValidation, "unknown block type", "unknown_block_type")
	InvalidBlockTitle      = errors.NewService(errors.ErrValidation, "invalid block title", "invalid_block_title")
	InvalidBlockGamesCount = errors.NewService(errors.ErrValidation, "invalid block games count", "invalid_block_games_count")

	// SocialLinks
	SocialLinkIncorrectURL = errors.NewService(errors.ErrValidation, "social link incorrect url", "social_link_incorrect_url")

	// Localization service
	LocalizationNotFound            = errors.NewService(errors.ErrNotFound, "localization not found", "localization_not_found")
	InvalidLocalizationLanguageCode = errors.NewService(errors.ErrValidation, "invalid localization language code", "invalid_localization_language_code")

	// Rating service
	RatingNotFound        = errors.NewService(errors.ErrNotFound, "rating not found", "rating_not_found")
	RatingUndefinedAgency = errors.NewService(errors.ErrNotFound, "undefined agency", "rating_undefined_agency")
	RatingUndefinedRating = errors.NewService(errors.ErrNotFound, "undefined rating", "rating_undefined_rating")

	// Review service
	ReviewMax3Available = errors.NewService(errors.ErrValidation, "maximum 3 reviews available", "maximum_3_reviews_available")

	// Document service
	DocumentNotFound            = errors.NewService(errors.ErrNotFound, "document not found", "document_not_found")
	DocumentAlreadyActivated    = errors.NewService(errors.ErrValidation, "document already activated", "document_already_activated")
	DocumentUnsupportedLanguage = errors.NewService(errors.ErrValidation, "unsupported document language", "document_unsupported_language")
)

func NewInternal(err error) errors.Error {
	return errors.NewInternal(err)
}

func NewValidation(err error) error {
	return errors.NewService(errors.ErrValidation, err.Error(), "validation")
}
