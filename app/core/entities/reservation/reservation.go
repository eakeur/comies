package reservation

import "comies/app/sdk/types"

type (
	Reservation struct {
		ID        types.ID
		ProductID types.ID
		Quantity  types.Quantity
		Ignore    []types.ID
		Replace   map[types.ID]types.ID
	}

	Failure struct {
		ProductID types.ID
		Error     error
	}
)
