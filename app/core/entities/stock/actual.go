package stock

import (
	"gomies/app/core/types/id"
	"gomies/app/core/types/quantity"
	"time"
)

// Actual is a wrapper for information about the stock computation
type Actual struct {
	// TargetID is the id of the object this computation is referencing to
	TargetID id.External

	// Actual is the quantity available in this stock
	Actual quantity.Quantity

	// InitialDate is the first date that counts for this stock computation
	InitialDate time.Time

	// FinalDate is the last date that counts for this stock computation
	FinalDate time.Time
}
