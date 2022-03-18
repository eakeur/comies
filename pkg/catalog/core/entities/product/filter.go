package product

import "gomies/pkg/sdk/types"

type (
	Filter struct {
		Code       string
		Name       string
		Type       Type
		CategoryID types.UID
		types.Filter
	}

	Key struct {
		ID   types.UID
		Code string
	}
)
