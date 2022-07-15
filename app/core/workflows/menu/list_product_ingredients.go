package menu

import (
	"comies/app/core/entities/ingredient"
	"comies/app/sdk/throw"
	"comies/app/sdk/types"
	"context"
)

func (w workflow) ListProductIngredients(ctx context.Context, productID types.ID) ([]ingredient.Ingredient, error) {

	if productID.Empty() {
		return []ingredient.Ingredient{}, throw.Error(throw.ErrMissingID)
	}

	list, err := w.ingredients.List(ctx, productID)
	if err != nil {
		return []ingredient.Ingredient{}, throw.Error(err).Params(map[string]interface{}{
			"product_id": productID.String(),
		})
	}

	return list, nil
}
