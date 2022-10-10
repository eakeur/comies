package menu

import (
	"comies/app/core/product"
	"comies/app/data/products"
	"context"
)

func ListProductsRunningOut(ctx context.Context) ([]product.Product, error) {
	list, err := products.ListRunningOut(ctx)
	if err != nil {
		return []product.Product{}, err
	}

	return list, err
}
