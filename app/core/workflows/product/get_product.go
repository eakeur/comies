package product

import (
	"context"
	"gomies/app/core/entities/catalog/product"
	"gomies/pkg/sdk/fault"
)

func (w workflow) GetProduct(ctx context.Context, ext product.Key) (product.Product, error) {
	const operation = "Workflows.Product.GetProduct"

	prod, err := w.products.Get(ctx, ext)
	if err != nil {
		return product.Product{}, fault.Wrap(err, operation)
	}
	return prod, nil
}
