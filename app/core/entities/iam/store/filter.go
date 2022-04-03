package store

import "gomies/pkg/sdk/types"

type (
	Filter struct {
		ParentExternalID types.UID
		types.Filter
	}

	Key struct {
		ID       types.UID
		Nickname string
	}
)
