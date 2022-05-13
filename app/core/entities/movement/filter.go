package movement

import (
	"gomies/app/sdk/fault"
	"gomies/app/sdk/listing"
	"gomies/app/sdk/types"
	"time"
)

type (
	Filter struct {
		ResourceID  types.ID
		InitialDate time.Time
		FinalDate   time.Time
		listing.Filter
	}
)

func (f Filter) Validate() error {
	if f.ResourceID.Empty() {
		return fault.ErrMissingUID
	}

	if f.FinalDate.Before(f.InitialDate) {
		return ErrInvalidPeriod
	}

	return nil
}
