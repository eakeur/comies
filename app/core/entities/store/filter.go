package store

import (
	"gomies/app/sdk/listing"
	"gomies/app/sdk/types"
)

type (
	Filter struct {
		ParentID types.ID
		listing.Filter
	}

	Key struct {
		ID       types.ID
		Nickname string
	}
)
