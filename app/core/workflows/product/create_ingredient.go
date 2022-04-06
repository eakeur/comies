package product

import (
	"context"
	"gomies/app/core/entities/catalog/product"
	"gomies/pkg/sdk/fault"
)

func (w workflow) CreateIngredient(ctx context.Context, productKey product.Key, ingredient product.Ingredient) (product.Ingredient, error) {
	const operation = "Workflows.Product.CreateIngredient"

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
