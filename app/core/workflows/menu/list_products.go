package menu

import (
	"comies/app/core/entities/product"
	"context"
)

func (w workflow) ListProducts(ctx context.Context, filter product.Filter) ([]product.Product, error) {

	list, err := w.products.List(ctx, filter)
	if err != nil {
		return []product.Product{}, err
	}

	return list, err
}
