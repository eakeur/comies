package menu

import (
	"comies/app/core/id"
	"time"
)

type (
	ProductFilter struct {
		Type       Type
		Code, Name string
	}

	MovementFilter struct {
		ProductID   id.ID
		InitialDate, FinalDate time.Time
	}
)


func ValidateMovementFilter(f MovementFilter) error {
	if err := id.ValidateID(f.ProductID); err != nil {
		return err
	}

	if f.FinalDate.Before(f.InitialDate) {
		return ErrInvalidPeriod
	}

	return nil
}