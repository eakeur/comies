package stock

import (
	"gomies/app/sdk/types"
	"time"
)

type (
	MovementType int

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
