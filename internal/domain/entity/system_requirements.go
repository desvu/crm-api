package entity

type SystemRequirements struct {
	Platform    uint
	Minimal     *RequirementsSet
	Recommended *RequirementsSet
}

type RequirementsSet struct {
	CPU       string
	GPU       string
	DiskSpace uint
	RAM       uint
}
