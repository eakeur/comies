package category

import (
	"gomies/app/sdk/listing"
	"gomies/app/sdk/types"
)

type (
	Filter struct {
		Code string
		Name string
		listing.Filter
	}

	Key struct {
		ID   types.ID
		Code string
		types.Store
	}
)
