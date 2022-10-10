package menu

import (
	"comies/app/core/product"
	"comies/app/data/products"
	"context"
)

func GetProductByCode(ctx context.Context, code string) (product.Product, error) {

	prod, err := products.GetByCode(ctx, code)
	if err != nil {
		return product.Product{}, err
	}
	return prod, nil
}
