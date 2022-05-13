package product

import (
	"gomies/app/sdk/listing"
	"gomies/app/sdk/types"
)

type (
	Filter struct {
		Code       string
		Name       string
		Type       Type
		CategoryID types.ID
		listing.Filter
	}

	Key struct {
		ID   types.ID
		Code string
	}
)
