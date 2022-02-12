package product

import (
	"context"
	"gomies/app/core/entities/product"
	"gomies/app/core/types/fault"
	"gomies/app/core/types/id"
)

func (w workflow) Get(ctx context.Context, ext id.External) (product.Product, error) {
	const operation = "Workflows.Product.Get"
	prod, err := w.products.Get(ctx, ext, product.All)
	if err != nil {
		return product.Product{}, fault.Wrap(err, operation)
	}
	return prod, nil
}
