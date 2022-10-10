package menu

import (
	"comies/app/core/product"
	"comies/app/data/products"
	"context"
)

func ListProducts(ctx context.Context, filter product.Filter) ([]product.Product, error) {

	list, err := products.List(ctx, filter)
	if err != nil {
		return []product.Product{}, err
	}

	return list, err
}
