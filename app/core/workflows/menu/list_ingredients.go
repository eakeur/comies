package menu

import (
	"comies/app/core/id"
	"comies/app/core/ingredient"
	"comies/app/core/types"
	"comies/app/data/ingredients"
	"context"
)

func ListIngredients(ctx context.Context, productID id.ID) ([]ingredient.Ingredient, error) {

	if productID.Empty() {
		return []ingredient.Ingredient{}, types.ErrMissingID
	}

	list, err := ingredients.List(ctx, productID)
	if err != nil {
		return []ingredient.Ingredient{}, err
	}

	return list, nil
}
