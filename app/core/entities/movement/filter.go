package movement

import (
	"comies/app/sdk/throw"
	"comies/app/sdk/types"
	"time"
)

type (
	Filter struct {
		ResourceID  types.ID
		InitialDate time.Time
		FinalDate   time.Time
	}
)

func (f Filter) Validate() error {
	if f.ResourceID.Empty() {
		return throw.ErrMissingID
	}

	if f.FinalDate.Before(f.InitialDate) {
		return ErrInvalidPeriod
	}

	return nil
}
