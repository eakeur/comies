package menu

import (
	"comies/app/core/types"
	"comies/app/data/ingredients"
	"context"
)

func RemoveIngredient(ctx context.Context, ingredientID types.ID) error {
	if err := types.ValidateID(ingredientID); err != nil {
		return err
	}

	err := ingredients.Remove(ctx, ingredientID)
	if err != nil {
		return err
	}

	return nil
}
