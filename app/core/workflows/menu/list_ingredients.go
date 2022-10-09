package menu

import (
	"comies/app/core/entities/ingredient"
	"comies/app/core/throw"
	"comies/app/core/types"
	"context"
)

func (w workflow) ListIngredients(ctx context.Context, productID types.ID) ([]ingredient.Ingredient, error) {

	if productID.Empty() {
		return []ingredient.Ingredient{}, throw.ErrMissingID
	}

	list, err := w.ingredients.List(ctx, productID)
	if err != nil {
		return []ingredient.Ingredient{}, err
	}

	return list, nil
}
