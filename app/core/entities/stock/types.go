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
		Count types.Quantity
	}

	Config struct {
		// MaximumQuantity is how many unities of this resource the stock can support
		MaximumQuantity types.Quantity

		// MinimumQuantity is the lowest quantity of this resource the stock can have
		MinimumQuantity types.Quantity
	}
)
