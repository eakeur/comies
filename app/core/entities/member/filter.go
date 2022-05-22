package member

import (
	"gomies/app/sdk/listing"
	"gomies/app/sdk/types"
)

type (
	Filter struct {
		Name string
		listing.Filter
	}

	Key struct {
		ID       types.ID
		Nickname string
	}
)
