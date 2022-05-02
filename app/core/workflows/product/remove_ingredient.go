package product

import (
	"context"
	"gomies/app/core/entities/catalog/product"
	"gomies/app/sdk/fault"
	"gomies/app/sdk/types"
)

func (w workflow) RemoveIngredient(ctx context.Context, productKey product.Key, id types.ID) error {
	const operation = "Workflows.Product.RemoveIngredient"

	err := w.products.RemoveIngredient(ctx, productKey, id)
	if err != nil {
		return fault.Wrap(err, operation)
	}
	return nil
}
