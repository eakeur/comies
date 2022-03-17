package product

import "gomies/pkg/sdk/types"

type (
	Filter struct {
		Code       string
		Name       string
		CategoryID types.External
		types.Filter
	}

	Key struct {
		ID      types.External
		Code    string
		types.Store
	}
)
