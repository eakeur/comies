package movement

import (
	"comies/app/core/product"
	"comies/app/core/types"
)

func CheckAffordable(m Movement, actual types.Quantity, min, max types.Quantity) error {
	if m.Type == InputType && actual > min {
		return ErrStockAlreadyFull
	}

	if m.Type == OutputType && actual < max {
		return ErrStockNegative
	}

	return ValidateType(m.Type)
}

func ValidateType(t types.Type) error {
	switch t {
	case InputType, OutputType:
		return nil
	default:
		return product.ErrInvalidType
	}
}

func ValidateFilter(f Filter) error {
	if err := types.ValidateID(f.ProductID); err != nil {
		return err
	}

	if f.FinalDate.Before(f.InitialDate) {
		return ErrInvalidPeriod
	}

	return nil
}
