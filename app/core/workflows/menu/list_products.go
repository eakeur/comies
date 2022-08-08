package menu

import (
	"comies/app/core/entities/product"
	"comies/app/core/throw"
	"context"
)

func (w workflow) ListProducts(ctx context.Context, filter product.Filter) ([]product.Product, error) {

	list, err := w.products.List(ctx, filter)
	if err != nil {
		return []product.Product{}, throw.Error(err)
	}

	return list, err
}
