package category

import (
	"gomies/pkg/sdk/listing"
	"gomies/pkg/sdk/types"
)

type (
	Filter struct {
		Code string
		Name string
		listing.Filter
	}

	Key struct {
		ID   types.UID
		Code string
		types.Store
	}
)
