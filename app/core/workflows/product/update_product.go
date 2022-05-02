package product

import (
	"context"
	"gomies/app/core/entities/catalog/product"
	"gomies/app/sdk/fault"
)

func (w workflow) UpdateProduct(ctx context.Context, prd product.Product) error {
	const operation = "Workflows.Product.UpdateProduct"

	if err := prd.Validate(); err != nil {
		return fault.Wrap(err, operation)
	}

	err := w.products.UpdateProduct(ctx, prd)
	if err != nil {
		return fault.Wrap(err, operation)
	}

	return nil
}
