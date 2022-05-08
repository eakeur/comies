package product

import "gomies/app/sdk/types"

type (
	Ingredient struct {
		ID           types.ID
		Quantity     types.Quantity
		ProductID    types.ID
		IngredientID types.ID
		Optional     bool
		Active       bool
		History      types.History
		types.Store
	}
)

func (i Ingredient) Validate() error {
	if i.IngredientID.Empty() {
		return ErrInvalidIngredient
	}
	if i.ProductID.Empty() {
		return ErrInvalidIngredient
	}

	if i.Quantity <= 0 {
		return ErrInvalidIngredientQuantity
	}

	return nil
}
