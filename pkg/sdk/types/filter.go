package types

type SortDirection string

const (
	DescendingSort SortDirection = "DESC"
	AscendingSort  SortDirection = "ASC"
)

type Filter struct {
	SortBy        string
	RangeStart    int
	RangeEnd      int
	SortDirection SortDirection
	Store
}
