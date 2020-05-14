package storefront

import (
	"strconv"
	"time"

	"github.com/qilin/crm-api/internal/domain/entity"
	"github.com/qilin/crm-api/internal/domain/enum/block"
)

//swagger:model StorefrontList
type storefrontList []storefrontInfo

//swagger:model StorefrontInfo
type storefrontInfo struct {
	// read-only: true
	// example: 12
	ID string `json:"id"`

	// example: Black-Friday Sales
	// required: true
	Name string `json:"name" validate:"required"`

	// read-only: true
	// example: 4
	Version uint `json:"version"`

	// read-only: true
	IsActive bool `json:"is_active"`

	// read-only: true
	CreatedAt time.Time `json:"created_at"`

	// read-only: true
	UpdatedAt time.Time `json:"updated_at"`
}

//swagger:model Storefront
type storefront struct {
	storefrontInfo
	Blocks []sfblock `json:"blocks"`
}

// Block is storefront template page block
//
// swagger: model StorefrontBlock
type sfblock struct {
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

func (s *storefront) GetBlocks() []entity.Block {
	var result []entity.Block
	for i := range s.Blocks {
		result = append(result, entity.Block(s.Blocks[i]))
	}
	return result
}

func (h *Handler) view(sf *entity.Storefront) *storefront {
	return &storefront{
		storefrontInfo: h.viewInfo(sf),
		Blocks:         h.viewBlocks(sf.Blocks),
	}
}

func (h *Handler) viewBlocks(b []entity.Block) []sfblock {
	var result []sfblock
	for i := range b {
		result = append(result, sfblock(b[i]))

	}
	return result
}

func (h *Handler) viewInfo(sf *entity.Storefront) storefrontInfo {
	return storefrontInfo{
		ID:        strconv.FormatUint(uint64(sf.ID), 10),
		Name:      sf.Name,
		IsActive:  sf.IsActive,
		Version:   sf.Version,
		CreatedAt: sf.CreatedAt,
		UpdatedAt: sf.UpdatedAt,
	}
}

func (h *Handler) viewList(sf []*entity.Storefront) storefrontList {
	var res = make([]storefrontInfo, len(sf))
	for i := range sf {
		res[i] = h.viewInfo(sf[i])
	}
	return res
}
