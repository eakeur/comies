package product

import (
	"context"
	"gomies/pkg/menu/core/entities/product"
	"gomies/pkg/sdk/fault"
	"gomies/pkg/sdk/session"
)

func (w workflow) List(ctx context.Context, filter product.Filter) ([]product.Product, error) {
	const operation = "Workflows.Product.List"

	_, err := session.DelegateSessionProps(ctx, operation, &filter.Store, nil)
	if err != nil {
		return []product.Product{}, fault.Wrap(err, operation)
	}

	list, err := w.products.List(ctx, filter)
	if err != nil {
		return []product.Product{}, fault.Wrap(err, operation)
	}
	return list, err
}
