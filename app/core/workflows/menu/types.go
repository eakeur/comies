package menu

import (
	"comies/app/core/entities/ingredient"
	"comies/app/sdk/types"
)

type (
	Reservation struct {
		ID        types.ID
		ProductID types.ID
		Quantity  types.Quantity
		Ignore    ingredient.IgnoredList
		Replace   ingredient.ReplacedList
		Failures  []ItemFailed
	}

	ItemFailed struct {
		ProductID types.ID
		Error     error
	}

	ActualBalance struct {
		ID    types.ID
		Count types.Quantity
	}
)
