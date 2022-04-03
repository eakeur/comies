package product

import "gomies/pkg/sdk/types"

type (
	Ingredient struct {
		types.Entity
		Quantity             types.Quantity
		ProductID            types.ID
		ProductExternalID    types.UID
		IngredientID         types.ID
		IngredientExternalID types.UID
		types.Store
	}
)

func (i Ingredient) Validate() error {
	if i.IngredientExternalID.Empty() {
		return ErrInvalidIngredient
	}

	if i.Quantity <= 0 {
		return ErrInvalidIngredientQuantity
	}

	return nil
}
