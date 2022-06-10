package member

import (
	"gomies/app/sdk/types"
)

type (
	Filter struct {
		Name string
	}

	Key struct {
		ID       types.ID
		Nickname string
	}
)
