package menu

import (
	"comies/app/core/ingredient"
	"comies/app/core/types"
	"comies/app/data/ingredients"
	"context"
)

func ListIngredients(ctx context.Context, productID types.ID) ([]ingredient.Ingredient, error) {
	if err := types.ValidateID(productID); err != nil {
		return nil, err
	}

	return ingredients.List(ctx, productID)
}
