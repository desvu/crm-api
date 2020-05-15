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
	License            string                `pg:"license,notnull,use_zero"`
	Trailer            string                `pg:"trailer,notnull,use_zero"`
	Status             uint8                 `pg:"status,notnull,use_zero"`
	PlayTime           uint                  `pg:"play_time,notnull,use_zero"`
	Platforms          []uint8               `pg:"platforms,array,notnull,use_zero"`
	ReleaseDate        time.Time             `pg:"release_date,notnull,use_zero"`
	PublishedAt        *time.Time            `pg:"published_at"`
	SocialLinks        *[]SocialLinks        `pg:"type:jsonb"`
	SystemRequirements *[]SystemRequirements `pg:"type:jsonb"`

	tableName struct{} `pg:"game_revisions"`
}

func (m model) Convert() *entity.GameRevision {
	return &entity.GameRevision{
		ID:                 m.ID,
		GameID:             m.GameID,
		Summary:            m.Summary,
		Description:        m.Description,
		License:            m.License,
		Trailer:            m.Trailer,
		Status:             game_revision.NewStatus(m.Status),
		Platforms:          game.NewPlatformArray(m.Platforms...),
		ReleaseDate:        m.ReleaseDate,
		PublishedAt:        m.PublishedAt,
		SocialLinks:        *convertSocialLinks(m.SocialLinks),
		SystemRequirements: *convertSystemRequirements(m.SystemRequirements),
		PlayTime:           m.PlayTime,
	}
}

func newModel(i *entity.GameRevision) (*model, error) {
	return &model{
		ID:                 i.ID,
		GameID:             i.GameID,
		Summary:            i.Summary,
		Description:        i.Description,
		License:            i.License,
		Trailer:            i.Trailer,
		Status:             i.Status.Value(),
		Platforms:          i.Platforms.Values(),
		ReleaseDate:        i.ReleaseDate,
		PublishedAt:        i.PublishedAt,
		SocialLinks:        newSocialLinksModel(&i.SocialLinks),
		SystemRequirements: newSystemRequirementsModel(&i.SystemRequirements),
		PlayTime:           i.PlayTime,
	}, nil
}

func newSocialLinksModel(i *[]entity.SocialLink) *[]SocialLinks {
	var socialLinks = make([]SocialLinks, 0)
	for _, item := range *i {
		socialLinks = append(socialLinks, SocialLinks{
			URL: item.URL,
		})
	}
	return &socialLinks
}

func convertSocialLinks(m *[]SocialLinks) *[]entity.SocialLink {
	var socialLinks = make([]entity.SocialLink, 0)
	for _, item := range *m {
		socialLinks = append(socialLinks, entity.SocialLink{
			URL: item.URL,
		})
	}
	return &socialLinks
}

func newSystemRequirementsModel(i *[]entity.SystemRequirements) *[]SystemRequirements {
	var systemRequirements = make([]SystemRequirements, 0)
	for _, item := range *i {
		requirementsSet := SystemRequirements{
			Platform: item.Platform.Value(),
			Minimal: RequirementsSetModel{
				CPU:       item.Minimal.CPU,
				GPU:       item.Minimal.GPU,
				DiskSpace: item.Minimal.DiskSpace,
				RAM:       item.Minimal.RAM,
			},
			Recommended: RequirementsSetModel{
				CPU:       item.Recommended.CPU,
				GPU:       item.Recommended.GPU,
				DiskSpace: item.Recommended.DiskSpace,
				RAM:       item.Recommended.RAM,
			},
		}

		systemRequirements = append(systemRequirements, requirementsSet)
	}
	return &systemRequirements
}

func convertSystemRequirements(m *[]SystemRequirements) *[]entity.SystemRequirements {
	var systemRequirements = make([]entity.SystemRequirements, 0)
	for _, item := range *m {
		requirementsSet := entity.SystemRequirements{
			Platform: game.NewPlatform(item.Platform),
			Minimal: &entity.RequirementsSet{
				CPU:       item.Minimal.CPU,
				GPU:       item.Minimal.GPU,
				DiskSpace: item.Minimal.DiskSpace,
				RAM:       item.Minimal.RAM,
			},
			Recommended: &entity.RequirementsSet{
				CPU:       item.Recommended.CPU,
				GPU:       item.Recommended.GPU,
				DiskSpace: item.Recommended.DiskSpace,
				RAM:       item.Recommended.RAM,
			},
		}

		systemRequirements = append(systemRequirements, requirementsSet)
	}
	return &systemRequirements
}

type SocialLinks struct {
	URL string `json:"url"`
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
