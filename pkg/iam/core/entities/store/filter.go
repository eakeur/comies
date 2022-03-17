package store

import "gomies/pkg/sdk/types"

type (
	Filter struct {
		ParentExternalID types.External
		types.Filter
	}

	Key struct {
		ID       types.External
		Nickname string
	}
)
