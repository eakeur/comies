package menu

import (
	"comies/app/core/menu/ingredient"
	"comies/app/core/types"
	"context"
)

func (w jobs) ListIngredients(ctx context.Context, productID types.ID) ([]ingredient.Ingredient, error) {

	if err := productID.Validate(); err != nil {
		return nil, err
	}

	return w.ingredients.ListByProductID(ctx, productID)
}
