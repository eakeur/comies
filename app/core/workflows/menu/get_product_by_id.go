package menu

import (
	"comies/app/core/id"
	"comies/app/core/product"
	"comies/app/data/products"
	"context"
)

func GetProductByID(ctx context.Context, id id.ID) (product.Product, error) {

	prod, err := products.GetByID(ctx, id)
	if err != nil {
		return product.Product{}, err
	}
	return prod, nil
}
