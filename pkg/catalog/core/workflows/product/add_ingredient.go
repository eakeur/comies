package product

import (
	"context"
	"gomies/pkg/catalog/core/entities/product"
	"gomies/pkg/sdk/fault"
)

func (w workflow) AddIngredient(ctx context.Context, productKey product.Key, ingredient product.Ingredient) (product.Ingredient, error) {
	const operation = "Workflows.Product.AddIngredient"

	if err := ingredient.Validate(); err != nil {
		return product.Ingredient{}, fault.Wrap(err, operation)
	}

	ingredient.ProductExternalID = productKey.ID
	_, err := w.products.SaveIngredients(ctx, productKey, ingredient)
	if err != nil {
		return product.Ingredient{}, fault.Wrap(err, operation)
	}

	return ingredient, nil
}
