package entity

import "github.com/qilin/crm-api/internal/domain/enum/game"

type SystemRequirements struct {
	Platform    game.Platform
	Minimal     *RequirementsSet
	Recommended *RequirementsSet
}

type RequirementsSet struct {
	CPU       string
	GPU       string
	DiskSpace uint
	RAM       uint
}
