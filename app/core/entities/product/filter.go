package product

import (
	"gomies/app/sdk/types"
)

type (
	Filter struct {
		Code       string
		Name       string
		Type       Type
		CategoryID types.ID
	}

	Key struct {
		ID   types.ID
		Code string
	}
)
