package storefront

import (
	"strconv"
	"time"

	"github.com/qilin/crm-api/internal/domain/entity"
)

type storefrontInfo struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Version   uint      `json:"version"`
	IsActive  bool      `json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

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

func (h *Handler) viewList(sf []*entity.Storefront) []storefrontInfo {
	var res = make([]storefrontInfo, len(sf))
	for i := range sf {
		res[i] = h.viewInfo(sf[i])
	}
	return res
}
