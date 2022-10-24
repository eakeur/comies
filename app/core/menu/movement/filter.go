package movement

import (
	"comies/app/core/types"
	"time"
)

type Filter struct {
	ProductID              types.ID
	InitialDate, FinalDate time.Time
}

func (f Filter) Validate() error {
	if err := f.ProductID.Validate(); err != nil {
		return err
	}

	if f.FinalDate.Before(f.InitialDate) {
		return ErrInvalidPeriod
	}

	return nil
}
