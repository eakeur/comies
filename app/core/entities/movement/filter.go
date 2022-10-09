package movement

import (
	"comies/app/core/types"
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
		return types.ErrMissingID
	}

	if f.FinalDate.Before(f.InitialDate) {
		return ErrInvalidPeriod
	}

	return nil
}
