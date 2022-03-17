package product

import (
	"context"
	"gomies/pkg/menu/core/entities/product"
	"gomies/pkg/sdk/fault"
	"gomies/pkg/sdk/session"
)

func (w workflow) Get(ctx context.Context, ext product.Key) (product.Product, error) {
	const operation = "Workflows.Product.Get"

	_, err := session.FromContext(ctx, operation)
	if err != nil {
		return product.Product{}, fault.Wrap(err, operation)
	}

	prod, err := w.products.Get(ctx, ext)
	if err != nil {
		return product.Product{}, fault.Wrap(err, operation)
	}
	return prod, nil
}
