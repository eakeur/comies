package product

import (
	"context"
	"gomies/app/core/entities/catalog/product"
	"gomies/pkg/sdk/fault"
)

func (w workflow) ListProducts(ctx context.Context, filter product.Filter) ([]product.Product, error) {
	const operation = "Workflows.Product.ListProducts"

	list, err := w.products.ListProducts(ctx, filter)
	if err != nil {
		return []product.Product{}, fault.Wrap(err, operation)
	}
	return list, err
}
