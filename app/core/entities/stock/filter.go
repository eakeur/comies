package stock

import (
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
		return ErrMissingResourceID
	}

	if f.FinalDate.Before(f.InitialDate) {
		return ErrInvalidPeriod
	}

	return nil
}
