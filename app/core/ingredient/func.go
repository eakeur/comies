package ingredient

import (
	"comies/app/core/types"
)

func Validate(i Ingredient) error {
	if err := types.ValidateID(i.IngredientID); err != nil {
		return ErrInvalidComponentID
	}

	if err := types.ValidateID(i.ProductID); err != nil {
		return ErrInvalidCompositeID
	}

	if i.Quantity <= 0 {
		return ErrInvalidQuantity
	}

	return nil
}
