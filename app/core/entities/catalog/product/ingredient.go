package product

import "gomies/app/sdk/types"

type (
	Ingredient struct {
		types.Entity
		Quantity     types.Quantity
		ProductID    types.ID
		IngredientID types.ID
		types.Store
	}
)

func (i Ingredient) Validate() error {
	if i.IngredientID.Empty() {
		return ErrInvalidIngredient
	}

	if i.Quantity <= 0 {
		return ErrInvalidIngredientQuantity
	}

	return nil
}
