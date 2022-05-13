package stock

import (
	"gomies/app/sdk/types"
)

type (
	Stock struct {
		ID      types.ID
		Active  bool
		History types.History
		// TargetID is an identifier for the object this stocks references to
		TargetID types.ID
		// MaximumQuantity is how many unities of this resource the stock can support
		MaximumQuantity types.Quantity
		// MinimumQuantity is the lowest quantity of this resource the stock can have
		MinimumQuantity types.Quantity
		// Location is a brief description of where this stock is located
		Location string
	}
)
