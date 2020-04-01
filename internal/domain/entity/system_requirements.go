package entity

type SystemRequirements struct {
	Platform    string
	Minimal     *RequirementsSet
	Recommended *RequirementsSet
}

type RequirementsSet struct {
	CPU       *string
	DiskSpace *string
	Gpu       *string
	Os        *string
	RAM       *string
}
