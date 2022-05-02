package product

import (
	"context"
	"gomies/app/core/entities/catalog/product"
	"gomies/app/sdk/fault"
)

func (w workflow) ListProducts(ctx context.Context, filter product.Filter) ([]product.Product, int, error) {
	const operation = "Workflows.Product.ListProducts"

	list, count, err := w.products.ListProducts(ctx, filter)
	if err != nil {
		return []product.Product{}, 0, fault.Wrap(err, operation)
	}
	return list, count, err
}
