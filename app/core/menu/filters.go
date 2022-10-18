package menu

import (
	"comies/app/core/types"
	"time"
)

type (
	ProductFilter struct {
		Type       Type
		Code, Name string
	}

	MovementFilter struct {
		ProductID              types.ID
		InitialDate, FinalDate time.Time
	}
)

func ValidateMovementFilter(f MovementFilter) error {
	if err := types.ValidateID(f.ProductID); err != nil {
		return err
	}

	if f.FinalDate.Before(f.InitialDate) {
		return ErrInvalidPeriod
	}

	return nil
}
