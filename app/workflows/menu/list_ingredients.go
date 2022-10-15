package menu

import (
	"comies/app/core/id"
	"comies/app/core/menu"
	"comies/app/data/ingredients"
	"context"
)

func ListIngredients(ctx context.Context, productID id.ID) ([]menu.Ingredient, error) {
	if err := id.ValidateID(productID); err != nil {
		return nil, err
	}

	return ingredients.List(ctx, productID)
}
