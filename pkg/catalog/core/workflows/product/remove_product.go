package product

import (
	"context"
	"gomies/pkg/catalog/core/entities/product"
	"gomies/pkg/sdk/fault"
)

func (w workflow) RemoveProduct(ctx context.Context, ext product.Key) error {
	const operation = "Workflows.Product.RemoveProduct"

	err := w.products.Remove(ctx, ext)
	if err != nil {
		return fault.Wrap(err, operation)
	}

	return err

}
