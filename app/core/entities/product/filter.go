package product

import "gomies/app/core/types/id"

type SortField int

const (
	Code         SortField = iota
	Name         SortField = iota
)

type Filter struct {
	Code       string
	Name       string
	CategoryID id.External
	SortBy     SortField
	Page       int
}
