package menu

import (
	"comies/app/core/id"
	"comies/app/core/product"
	"comies/app/data/products"
	"context"
)

func CreateProduct(ctx context.Context, prd product.Product) (product.Product, error) {

	id.Create(&prd.ID)

	if err := prd.Validate(); err != nil {
		return product.Product{}, err
	}

	prd, err := products.Create(ctx, prd)
	if err != nil {
		return product.Product{}, err
	}

	return prd, nil
}
