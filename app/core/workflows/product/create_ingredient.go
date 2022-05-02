package product

import (
	"context"
	"gomies/app/core/entities/catalog/product"
	"gomies/app/sdk/fault"
)

func (w workflow) CreateIngredient(ctx context.Context, productKey product.Key, input IngredientInput) (product.Ingredient, error) {
	const operation = "Workflows.Product.CreateIngredient"

	ingredient := product.Ingredient{
		ProductID:    productKey.ID,
		IngredientID: input.IngredientID,
		Quantity:     input.Quantity,
	}

	if err := ingredient.Validate(); err != nil {
		return product.Ingredient{}, fault.Wrap(err, operation)
	}

	_, err := w.products.SaveIngredients(ctx, productKey, ingredient)
	if err != nil {
		return product.Ingredient{}, fault.Wrap(err, operation)
	}

	return ingredient, nil
}
