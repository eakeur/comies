package menu

import (
	"context"
	"gomies/app/core/entities/product"
	"gomies/app/sdk/fault"
)

func (w workflow) RemoveProduct(ctx context.Context, ext product.Key) error {
	const operation = "Workflows.Product.RemoveProduct"

	err := w.products.RemoveProduct(ctx, ext)
	if err != nil {
		return fault.Wrap(err, operation)
	}

	return err

}
