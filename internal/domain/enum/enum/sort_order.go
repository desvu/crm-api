package enum

type SortOrderType uint8

const (
	SortOrderDesc SortOrderType = 1
	SortOrderAsc  SortOrderType = 2
)

type SortOrderColumn uint8

const (
	SortOrderColumnReleaseDate SortOrderColumn = 1
	SortOrderColumnName        SortOrderColumn = 2
)
