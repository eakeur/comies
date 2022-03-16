package category

import "gomies/app/core/types/filter"

type Filter struct {
	Code string
	Name string
	filter.Filter
}
