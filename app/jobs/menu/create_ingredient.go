package menu

import (
	"comies/app/core/menu/ingredient"
	"context"
)

func (w jobs) CreateIngredient(ctx context.Context, i ingredient.Ingredient) (ingredient.Ingredient, error) {
	save, err := i.WithID(w.createID()).Validate()
	if err != nil {
		return ingredient.Ingredient{}, err
	}

	prd, err := w.products.GetByID(ctx, i.ProductID)
	if err != nil || !prd.IsComposite() {
		return ingredient.Ingredient{}, ingredient.ErrInvalidCompositeID
	}

	ing, err := w.products.GetByID(ctx, i.IngredientID)
	if err != nil || ing.IsOutput() {
		return ingredient.Ingredient{}, ingredient.ErrInvalidComponentID
	}

	return i, w.ingredients.Create(ctx, save)
}
