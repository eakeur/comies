package menu

import (
	"comies/app/core/id"
	"comies/app/data/ingredients"
	"context"
)

func RemoveIngredient(ctx context.Context, ingredientID id.ID) error {
	if err := id.ValidateID(ingredientID); err != nil {
		return err
	}

	err := ingredients.Remove(ctx, ingredientID)
	if err != nil {
		return err
	}

	return nil
}
