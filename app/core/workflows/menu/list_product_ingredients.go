package menu

import (
	"context"
	"gomies/app/core/entities/ingredient"
	"gomies/app/sdk/fault"
	"gomies/app/sdk/types"
)

func (w workflow) ListProductIngredients(ctx context.Context, productID types.ID) ([]ingredient.Ingredient, error) {

	if productID.Empty() {
		return []ingredient.Ingredient{}, fault.Wrap(fault.ErrMissingID)
	}

	list, err := w.ingredients.List(ctx, productID)
	if err != nil {
		return []ingredient.Ingredient{}, fault.Wrap(err).Params(map[string]interface{}{
			"product_id": productID.String(),
		})
	}

	return list, nil
}
