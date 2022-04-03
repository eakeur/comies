package store

import (
	"gomies/pkg/sdk/listing"
	"gomies/pkg/sdk/types"
)

type (
	Filter struct {
		ParentExternalID types.UID
		listing.Filter
	}

	Key struct {
		ID       types.UID
		Nickname string
	}
)
