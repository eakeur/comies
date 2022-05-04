package stock

import "gomies/app/sdk/types"

type (
	Reservation struct {
		ResourceID types.ID
		Quantity   types.Quantity
	}

	ReservationResult struct {
		ResourceID types.ID
		Got        types.Quantity
		Want       types.Quantity
		Error      error
	}
)
