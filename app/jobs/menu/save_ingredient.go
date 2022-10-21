package menu

import (
	"comies/app/core/ingredient"
	"comies/app/core/product"
	"comies/app/data/ingredients"
	"comies/app/data/products"
	"context"
)

func ValidateIngredient(ctx context.Context, i ingredient.Ingredient) (ingredient.Ingredient, error) {
	prd, err := products.GetByID(ctx, i.ProductID)
	if err != nil || !product.IsComposite(prd.Type) {
		return ingredient.Ingredient{}, ingredient.ErrInvalidCompositeID
	}

	ing, err := products.GetByID(ctx, i.IngredientID)
	if err != nil || product.IsOutput(ing.Type) {
		return ingredient.Ingredient{}, ingredient.ErrInvalidComponentID
	}

	return i, ingredients.Create(ctx, i)
}
