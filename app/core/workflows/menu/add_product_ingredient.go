package menu

import (
	"comies/app/core/entities/ingredient"
	"comies/app/sdk/fault"
	"context"
)

func (w workflow) AddProductIngredient(ctx context.Context, i ingredient.Ingredient) (ingredient.Ingredient, error) {

	if err := i.Validate(); err != nil {
		return ingredient.Ingredient{}, fault.Wrap(err).Params(map[string]interface{}{
			"product_id":    i.ProductID,
			"ingredient_id": i.IngredientID,
			"quantity":      i.Quantity,
		})
	}

	i, err := w.ingredients.Create(ctx, i)
	if err != nil {
		return ingredient.Ingredient{}, fault.Wrap(err)
	}

	return i, nil
}
