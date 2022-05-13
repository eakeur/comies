package stocking

import (
	"gomies/app/sdk/types"
	"time"
)

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

	CloseRequest struct {
		// InitialDate is the first date that counts for this stock close
		InitialDate time.Time

		// FinalDate is the last date that counts for this stock close
		FinalDate time.Time
	}

	AdditionResult struct {
		ID    types.ID
		Count types.Quantity
	}
)
