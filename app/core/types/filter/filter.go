package filter

type SortDirection string

const (
	DescendingSort SortDirection = "DESC"
	AscendingSort  SortDirection = "ASC"
)

type Filter struct {
	SortBy        string
	SortDirection SortDirection
	RangeStart    int
	RangeEnd      int
}
