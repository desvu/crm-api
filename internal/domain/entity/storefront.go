package entity

import (
	"time"

	"github.com/qilin/crm-api/internal/domain/enum/block"
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
type Block struct {
	// Type is how block must be rendered
	Type block.Type `json:"type"`

	// Title is user facing title for page block
	Title block.Title `json:"title"`

	// Filter is query to catalog for 'view more' button
	Filter string `json:"filter"`

	// GameIDs is list of games identifiers
	GameIDs []string `json:"games"`
}
