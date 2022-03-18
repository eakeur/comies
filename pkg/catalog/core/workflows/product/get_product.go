package product

import (
	"context"
	"gomies/pkg/catalog/core/entities/product"
	"gomies/pkg/sdk/fault"
	"gomies/pkg/sdk/session"
)

func (w workflow) GetProduct(ctx context.Context, ext product.Key) (product.Product, error) {
	const operation = "Workflows.Product.GetProduct"

	_, err := session.DelegateSessionProps(ctx, operation, &ext.Store, nil)
	if err != nil {
		return product.Product{}, fault.Wrap(err, operation)
	}

	prod, err := w.products.Get(ctx, ext)
	if err != nil {
		return product.Product{}, fault.Wrap(err, operation)
	}
	return prod, nil
}
