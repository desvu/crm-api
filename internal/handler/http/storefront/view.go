package storefront

import (
	"strconv"
	"time"

	"github.com/qilin/crm-api/internal/domain/entity"
)

//swagger:model StorefrontList
type storefrontList []storefrontInfo

//swagger:model StorefrontInfo
type storefrontInfo struct {
	// read-only: true
	// example: 12
	ID string `json:"id"`

	// example: Black-Friday Sales
	Name string `json:"name"`

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
	Blocks []entity.Block `json:"blocks"`
}

func (h *Handler) view(sf *entity.Storefront) *storefront {
	return &storefront{
		storefrontInfo: h.viewInfo(sf),
		Blocks:         sf.Blocks,
	}
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
