package crew

import (
	"gomies/pkg/sdk/types"
)

type (
	Filter struct {
		Name string
		types.Filter
	}

	Key struct {
		ID       types.External
		Nickname string
		StoreID  types.External
	}
)
