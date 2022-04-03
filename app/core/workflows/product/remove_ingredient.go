package product

import (
	"context"
	"gomies/app/core/entities/catalog/product"
	"gomies/pkg/sdk/fault"
	"gomies/pkg/sdk/types"
)

func (w workflow) RemoveIngredient(ctx context.Context, productKey product.Key, id types.UID) error {
	const operation = "Workflows.Product.RemoveIngredient"

	err := w.products.RemoveIngredient(ctx, productKey, id)
	if err != nil {
		return fault.Wrap(err, operation)
	}
	return nil
}
