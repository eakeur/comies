package product

import "comies/app/core/id"

type (
	Filter struct {
		Code       string
		Name       string
		Type       Type
		CategoryID id.ID
	}

	Key struct {
		ID   id.ID
		Code string
	}
)
