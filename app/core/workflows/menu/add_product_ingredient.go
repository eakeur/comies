package menu

import (
	"comies/app/core/entities/ingredient"
	"comies/app/sdk/throw"
	"context"
)

func (w workflow) AddProductIngredient(ctx context.Context, i ingredient.Ingredient) (ingredient.Ingredient, error) {

	w.id.Create(&i.ID)

	if err := i.Validate(); err != nil {
		return ingredient.Ingredient{}, throw.Error(err).Params(map[string]interface{}{
			"product_id":    i.ProductID,
			"ingredient_id": i.IngredientID,
			"quantity":      i.Quantity,
		})
	}

	i, err := w.ingredients.Create(ctx, i)
	if err != nil {
		return ingredient.Ingredient{}, throw.Error(err)
	}

	return i, nil
}
