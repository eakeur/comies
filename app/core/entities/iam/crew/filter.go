package crew

import (
	"gomies/pkg/sdk/listing"
	"gomies/pkg/sdk/types"
)

type (
	Filter struct {
		Name string
		listing.Filter
	}

	Key struct {
		ID       types.UID
		Nickname string
	}
)
