package movement

import (
	"comies/app/sdk/throw"
	"comies/app/sdk/types"
	"time"
)

type (
	Filter struct {
		ProductID   types.ID
		InitialDate time.Time
		FinalDate   time.Time
	}
)

func (f Filter) Validate() error {
	if f.ProductID.Empty() {
		return throw.ErrMissingID
	}

	if f.FinalDate.Before(f.InitialDate) {
		return ErrInvalidPeriod
	}

	return nil
}
