package product

import (
	"gomies/app/core/types/filter"
	"gomies/app/core/types/id"
)

type Filter struct {
	Code       string
	Name       string
	CategoryID id.External
	filter.Filter
}
