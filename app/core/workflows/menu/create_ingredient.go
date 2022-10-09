package menu

import (
	"comies/app/core/entities/ingredient"
	"comies/app/core/entities/product"
	"context"
)

func (w workflow) CreateIngredient(ctx context.Context, i ingredient.Ingredient) (ingredient.Ingredient, error) {

	w.id.Create(&i.ID)

	if err := i.Validate(); err != nil {
		return ingredient.Ingredient{}, err
	}

	compositeProduct, err := w.products.GetByID(ctx, i.ProductID)
	if err != nil {
		return ingredient.Ingredient{}, err
	}

	if compositeProduct.Type == product.InputType || compositeProduct.Type == product.OutputType {
		return ingredient.Ingredient{}, ingredient.ErrInvalidCompositeType
	}

	ingredientProduct, err := w.products.GetByID(ctx, i.IngredientID)
	if err != nil {
		return ingredient.Ingredient{}, err
	}

	if ingredientProduct.Type == product.OutputType || ingredientProduct.Type == product.OutputCompositeType {
		return ingredient.Ingredient{}, ingredient.ErrInvalidIngredientType
	}

	i, err = w.ingredients.Create(ctx, i)
	if err != nil {
		return ingredient.Ingredient{}, err
	}

	return i, nil
}
