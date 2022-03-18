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
		ID       types.UID
		Nickname string
	}
)
