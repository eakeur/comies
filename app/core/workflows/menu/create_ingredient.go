package menu

import (
	"comies/app/core/id"
	"comies/app/core/ingredient"
	"comies/app/core/product"
	"comies/app/data/ingredients"
	"comies/app/data/products"
	"context"
)

func CreateIngredient(ctx context.Context, i ingredient.Ingredient) (ingredient.Ingredient, error) {

	id.Create(&i.ID)

	if err := i.Validate(); err != nil {
		return ingredient.Ingredient{}, err
	}

	compositeProduct, err := products.GetByID(ctx, i.ProductID)
	if err != nil {
		return ingredient.Ingredient{}, err
	}

	if compositeProduct.Type == product.InputType || compositeProduct.Type == product.OutputType {
		return ingredient.Ingredient{}, ingredient.ErrInvalidCompositeType
	}

	ingredientProduct, err := products.GetByID(ctx, i.IngredientID)
	if err != nil {
		return ingredient.Ingredient{}, err
	}

	if ingredientProduct.Type == product.OutputType || ingredientProduct.Type == product.OutputCompositeType {
		return ingredient.Ingredient{}, ingredient.ErrInvalidIngredientType
	}

	i, err = ingredients.Create(ctx, i)
	if err != nil {
		return ingredient.Ingredient{}, err
	}

	return i, nil
}
