package product

import (
	"context"
	"gomies/app/core/entities/product"
	"gomies/app/core/types/fault"
)

func (w workflow) List(ctx context.Context, filter product.Filter) ([]product.Product, error) {
	const operation = "Workflows.Product.List"
	list, err := w.products.List(ctx, filter)
	if err != nil {
		return []product.Product{}, fault.Wrap(err, operation)
	}
	return list, err
}
