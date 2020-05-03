package postgres

import (
	"time"

	"github.com/qilin/crm-api/internal/domain/entity"
)

type storefront struct {
	ID        uint      `pg:"id"`
	Name      string    `pg:"name"`
	IsActive  bool      `pg:"-"`
	CreatedAt time.Time `pg:"created_at,default:now()"`

	// relations
	Version version
	// Versions    []version
	// Activation  *activation
	// Activations []activation

	tableName struct{} `pg:"storefronts,alias:sf"`
}

type version struct {
	StorefrontID uint           `pg:"storefront_id,pk,fk"`
	ID           uint           `pg:"id,use_zero,pk"`
	Blocks       []entity.Block `pg:"blocks"`
	CreatedAt    time.Time      `pg:"created_at,default:now()"`

	// relations
	// Activation *activation

	tableName struct{} `pg:"storefront_versions"`
}

type activation struct {
	Timestamp    time.Time `pg:"timestamp,default:now(),pk"`
	VersionID    uint      `pg:"version_id,use_zero,fk"`
	StorefrontID uint      `pg:"storefront_id,fk"`

	tableName struct{} `pg:"storefront_activations"`
}

func newStorefront(i *entity.Storefront) (*storefront, error) {
	return &storefront{
		ID:        i.ID,
		Name:      i.Name,
		CreatedAt: i.CreatedAt,
		Version: version{
			StorefrontID: i.ID,
			ID:           i.Version,
			Blocks:       i.Blocks,
		},
	}, nil
}

func (s *storefront) Convert() *entity.Storefront {
	return &entity.Storefront{
		ID:        s.ID,
		Name:      s.Name,
		IsActive:  s.IsActive,
		CreatedAt: s.CreatedAt,
		Version:   s.Version.ID,
		Blocks:    s.Version.Blocks,
		UpdatedAt: s.Version.CreatedAt,
	}
}
