package product

import (
	"context"
	"gomies/app/core/entities/catalog/product"
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
