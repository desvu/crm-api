package entity

import (
	"time"

	"github.com/qilin/crm-api/internal/domain/enum/block"
	"github.com/qilin/crm-api/internal/domain/errors"
)

type Storefront struct {
	// ID is unique identifier for storefront page template
	ID uint

	// Name is user-friendly template name
	Name string

	// IsActive is status flag
	IsActive bool

	// Version is latest template version number
	Version uint

	// Blocks are content of page
	Blocks []Block

	// CreatedAt is template creation date
	CreatedAt time.Time

	// UpdatedAt is last update date
	UpdatedAt time.Time
}

// Block is storefront template page block
//
// swagger: model
type Block struct {
	// Type is how block must be rendered
	//
	// required: true
	// example: basic_layout_2
	Type block.Type `json:"type"`

	// Title is user facing title for page block
	//
	// example: special-offer
	Title block.Title `json:"title"`

	// Filter is query to catalog for 'view more' button
	//
	// example: /catalog?genre=racing
	Filter string `json:"filter"`

	// GameIDs is list of games identifiers
	//
	// example: ["c6afe465-f6c0-46aa-84f2-471ab2280960","a62a08fe-457f-480c-a417-f1e872485df9","401b8160-1d71-47c2-b232-7e05c129ef15","e31c1a39-143b-49db-87cf-b4525beacfdd"]
	GameIDs []string `json:"games"`
}

func (b *Block) Validate() error {
	if !b.Type.Valid() {
		return errors.UnknownBlockType
	}
	if !b.Title.Valid() {
		return errors.InvalidBlockTitle
	}
	if !b.Type.ValidGamesCount(len(b.GameIDs)) {
		return errors.InvalidBlockGamesCount
	}
	return nil
}
