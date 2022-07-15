package menu

import (
	"comies/app/core/entities/ingredient"
	"comies/app/sdk/fault"
	"comies/app/sdk/types"
	"context"
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
