package menu

import (
	"context"
	product2 "gomies/app/core/entities/product"
	"gomies/app/sdk/fault"
)

func (w workflow) ListProducts(ctx context.Context, filter product2.Filter) ([]product2.Product, int, error) {

	list, count, err := w.products.ListProducts(ctx, filter)
	if err != nil {
		return []product2.Product{}, 0, fault.Wrap(err)
	}
	return list, count, err
}
