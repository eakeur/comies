package menu

import (
	"comies/app/core/entities/ingredient"
	"comies/app/core/entities/product"
	"comies/app/core/throw"
	"context"
)

func (w workflow) CreateIngredient(ctx context.Context, i ingredient.Ingredient) (ingredient.Ingredient, error) {

	var params = map[string]interface{}{
		"product_id":    i.ProductID,
		"ingredient_id": i.IngredientID,
		"quantity":      i.Quantity,
	}

	w.id.Create(&i.ID)

	if err := i.Validate(); err != nil {
		return ingredient.Ingredient{}, throw.Error(err).Params(params)
	}

	compositeProduct, err := w.products.GetByID(ctx, i.ProductID)
	if err != nil {
		return ingredient.Ingredient{}, throw.Error(err).Params(params)
	}

	if compositeProduct.Type == product.InputType {
		return ingredient.Ingredient{}, ingredient.ErrInvalidCompositeType
	}

	ingredientProduct, err := w.products.GetByID(ctx, i.IngredientID)
	if err != nil {
		return ingredient.Ingredient{}, throw.Error(err).Params(params)
	}

	if ingredientProduct.Type == product.OutputType || ingredientProduct.Type == product.OutputCompositeType {
		return ingredient.Ingredient{}, ingredient.ErrInvalidIngredientType
	}

	i, err = w.ingredients.Create(ctx, i)
	if err != nil {
		return ingredient.Ingredient{}, throw.Error(err)
	}

	return i, nil
}
