package postgres

import (
	"time"

	"github.com/qilin/crm-api/internal/domain/entity"
	"github.com/qilin/crm-api/internal/domain/enum/game"
	"github.com/qilin/crm-api/internal/domain/enum/game_revision"
)

type model struct {
	ID                 uint                  `pg:"id"`
	GameID             string                `pg:"game_id,notnull,use_zero"`
	Summary            string                `pg:"summary,notnull,use_zero"`
	Description        string                `pg:"description,notnull,use_zero"`
	Slug               string                `pg:"slug,notnull,use_zero"`
	License            string                `pg:"license,notnull,use_zero"`
	Status             uint8                 `pg:"status,notnull,use_zero"`
	Platforms          []uint8               `pg:"platforms,array,notnull,use_zero"`
	ReleaseDate        time.Time             `pg:"release_date,notnull,use_zero"`
	PublishedAt        *time.Time            `pg:"published_at"`
	SystemRequirements *[]SystemRequirements `pg:"type:jsonb"`
	tableName          struct{}              `pg:"game_revisions"`
}

func (m model) Convert() *entity.GameRevision {
	a := &entity.GameRevision{
		ID:                 m.ID,
		GameID:             m.GameID,
		Summary:            m.Summary,
		Description:        m.Description,
		Slug:               m.Slug,
		License:            m.License,
		Status:             game_revision.NewStatus(m.Status),
		Platforms:          game.NewPlatformArray(m.Platforms...),
		ReleaseDate:        m.ReleaseDate,
		PublishedAt:        m.PublishedAt,
		SystemRequirements: *convertSystemRequirements(m.SystemRequirements),
	}
	return a
}

func newModel(i *entity.GameRevision) (*model, error) {
	return &model{
		ID:                 i.ID,
		GameID:             i.GameID,
		Summary:            i.Summary,
		Description:        i.Description,
		Slug:               i.Slug,
		License:            i.License,
		Status:             i.Status.Value(),
		Platforms:          i.Platforms.Values(),
		ReleaseDate:        i.ReleaseDate,
		PublishedAt:        i.PublishedAt,
		SystemRequirements: newSystemRequirementsModel(&i.SystemRequirements),
	}, nil
}

func newSystemRequirementsModel(i *[]entity.SystemRequirements) *[]SystemRequirements {
	sysReqArray := []SystemRequirements{}
	for _, req := range *i {
		sysReq := SystemRequirements{
			Platform: req.Platform.Value(),
			Minimal: RequirementsSetModel{
				CPU:       req.Minimal.CPU,
				GPU:       req.Minimal.GPU,
				DiskSpace: req.Minimal.DiskSpace,
				RAM:       req.Minimal.RAM,
			},
			Recommended: RequirementsSetModel{
				CPU:       req.Recommended.CPU,
				GPU:       req.Recommended.GPU,
				DiskSpace: req.Recommended.DiskSpace,
				RAM:       req.Recommended.RAM,
			},
		}

		sysReqArray = append(sysReqArray, sysReq)
	}
	return &sysReqArray
}

func convertSystemRequirements(m *[]SystemRequirements) *[]entity.SystemRequirements {
	sysReqArray := []entity.SystemRequirements{}
	for _, req := range *m {
		sysReq := entity.SystemRequirements{
			Platform: game.NewPlatform(req.Platform),
			Minimal: &entity.RequirementsSet{
				CPU:       req.Minimal.CPU,
				GPU:       req.Minimal.GPU,
				DiskSpace: req.Minimal.DiskSpace,
				RAM:       req.Minimal.RAM,
			},
			Recommended: &entity.RequirementsSet{
				CPU:       req.Recommended.CPU,
				GPU:       req.Recommended.GPU,
				DiskSpace: req.Recommended.DiskSpace,
				RAM:       req.Recommended.RAM,
			},
		}

		sysReqArray = append(sysReqArray, sysReq)
	}
	return &sysReqArray
}

type SystemRequirements struct {
	Platform    uint8                `json:"platform"`
	Minimal     RequirementsSetModel `json:"minimal"`
	Recommended RequirementsSetModel `json:"recommended"`
}

type RequirementsSetModel struct {
	CPU       string `json:"cpu"`
	GPU       string `json:"gpu"`
	DiskSpace uint   `json:"disk_space"`
	RAM       uint   `json:"ram"`
}
