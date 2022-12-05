package menu

import (
	"comies/core/menu/ingredient"
	"comies/core/types"
	"context"
)

func (w jobs) ListIngredients(ctx context.Context, productID types.ID) ([]ingredient.Ingredient, error) {

	if err := productID.Validate(); err != nil {
		return nil, err
	}

	return w.ingredients.ListByProductID(ctx, productID)
}
