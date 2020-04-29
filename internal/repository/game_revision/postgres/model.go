package postgres

import (
	"time"

	"github.com/qilin/crm-api/internal/domain/entity"
	"github.com/qilin/crm-api/internal/domain/enum/game"
	"github.com/qilin/crm-api/internal/domain/enum/game_revision"
)

type model struct {
	ID                 uint                 `pg:"id"`
	GameID             string               `pg:"game_id,notnull,use_zero"`
	Summary            string               `pg:"summary,notnull,use_zero"`
	Description        string               `pg:"description,notnull,use_zero"`
	Slug               string               `pg:"slug,notnull,use_zero"`
	License            string               `pg:"license,notnull,use_zero"`
	Status             uint8                `pg:"status,notnull,use_zero"`
	Platforms          []uint8              `pg:"platforms,array,notnull,use_zero"`
	ReleaseDate        time.Time            `pg:"release_date,notnull,use_zero"`
	PublishedAt        *time.Time           `pg:"published_at"`
	SystemRequirements []systemRequirements `pg:system_requirements`
	tableName          struct{}             `pg:"game_revisions"`
}

func (m model) Convert() *entity.GameRevision {
	return &entity.GameRevision{
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
		SystemRequirements: convertSystemRequirements(m.SystemRequirements),
	}
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
		SystemRequirements: newSystemRequirementsModel(i.SystemRequirements),
	}, nil
}

func newSystemRequirementsModel(i []entity.SystemRequirements) []systemRequirements {
	sysReqArray := []systemRequirements{}
	for _, req := range i {
		sysReq := systemRequirements{
			Platform: req.Platform,
		}
		if req.Minimal != nil {
			sysReq.Minimal = &requirementsSetModel{
				CPU:       req.Minimal.CPU,
				GPU:       req.Minimal.GPU,
				DiskSpace: req.Minimal.DiskSpace,
				RAM:       req.Minimal.RAM,
			}
		}
		if req.Recommended != nil {
			sysReq.Minimal = &requirementsSetModel{
				CPU:       req.Recommended.CPU,
				GPU:       req.Recommended.GPU,
				DiskSpace: req.Recommended.DiskSpace,
				RAM:       req.Recommended.RAM,
			}
		}
		sysReqArray = append(sysReqArray)
	}
	return sysReqArray
}

func convertSystemRequirements(m []systemRequirements) []entity.SystemRequirements {
	sysReqArray := []entity.SystemRequirements{}
	for _, req := range m {
		sysReq := entity.SystemRequirements{
			Platform: req.Platform,
		}
		if req.Minimal != nil {
			sysReq.Minimal = &entity.RequirementsSet{
				CPU:       req.Minimal.CPU,
				GPU:       req.Minimal.GPU,
				DiskSpace: req.Minimal.DiskSpace,
				RAM:       req.Minimal.RAM,
			}
		}
		if req.Recommended != nil {
			sysReq.Recommended = &entity.RequirementsSet{
				CPU:       req.Recommended.CPU,
				GPU:       req.Recommended.GPU,
				DiskSpace: req.Recommended.DiskSpace,
				RAM:       req.Recommended.RAM,
			}
		}
		sysReqArray = append(sysReqArray, sysReq)
	}
	return sysReqArray
}

type systemRequirements struct {
	Platform    uint
	Minimal     *requirementsSetModel
	Recommended *requirementsSetModel
}

type requirementsSetModel struct {
	CPU       string
	GPU       string
	DiskSpace uint
	RAM       uint
}
