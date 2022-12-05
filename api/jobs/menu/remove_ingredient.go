package menu

import (
	"comies/core/types"
	"context"
)

func (w jobs) RemoveIngredient(ctx context.Context, productID, ingredientID types.ID) error {

	if err := productID.Validate(); err != nil {
		return err
	}

	if err := ingredientID.Validate(); err != nil {
		return err
	}

	return w.ingredients.Remove(ctx, productID, ingredientID)
}
