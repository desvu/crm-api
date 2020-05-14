package postgres

import (
	"time"

	"github.com/qilin/crm-api/internal/domain/entity"
	"github.com/qilin/crm-api/internal/domain/enum/block"
)

type storefront struct {
	ID        uint      `pg:"id"`
	Name      string    `pg:"name,use_zero"`
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
	StorefrontID uint      `pg:"storefront_id,pk,fk"`
	ID           uint      `pg:"id,use_zero,pk"`
	Blocks       []sfblock `pg:"blocks,use_zero"`
	CreatedAt    time.Time `pg:"created_at,default:now()"`

	// relations
	// Activation *activation

	tableName struct{} `pg:"storefront_versions"`
}

type sfblock struct {
	Type    block.Type  `json:"type"`
	Title   block.Title `json:"title"`
	Filter  string      `json:"filter"`
	GameIDs []string    `json:"games"`
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
			Blocks:       modelBlocks(i.Blocks),
		},
	}, nil
}

func modelBlocks(b []entity.Block) []sfblock {
	var res []sfblock
	for i := range b {
		res = append(res, sfblock(b[i]))
	}
	return res
}

func (s *storefront) Convert() *entity.Storefront {
	return &entity.Storefront{
		ID:        s.ID,
		Name:      s.Name,
		IsActive:  s.IsActive,
		CreatedAt: s.CreatedAt,
		Version:   s.Version.ID,
		Blocks:    s.convertBlocks(),
		UpdatedAt: s.Version.CreatedAt,
	}
}

func (s *storefront) convertBlocks() []entity.Block {
	var res []entity.Block
	for i := range s.Version.Blocks {
		res = append(res, entity.Block(s.Version.Blocks[i]))
	}
	return res
}
