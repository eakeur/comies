package stock

import (
	"gomies/pkg/sdk/types"
	"time"
)

// Actual is a wrapper for information about the stock computation
type Actual struct {
	// TargetID is the id of the object this computation is referencing to
	TargetID types.External

	// Actual is the quantity available in this stock
	Actual types.Quantity

	// InitialDate is the first date that counts for this stock computation
	InitialDate time.Time

	// FinalDate is the last date that counts for this stock computation
	FinalDate time.Time
}
