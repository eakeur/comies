package category

import "gomies/pkg/sdk/types"

type (
	Filter struct {
		Code string
		Name string
		types.Filter
	}

	Key struct {
		ID   types.External
		Code string
		types.Store
	}
)
