package menu

import (
	"comies/app/core/entities/product"
	"context"
)

func (w workflow) CreateProduct(ctx context.Context, prd product.Product) (product.Product, error) {

	w.id.Create(&prd.ID)

	if err := prd.Validate(); err != nil {
		return product.Product{}, err
	}

	prd, err := w.products.Create(ctx, prd)
	if err != nil {
		return product.Product{}, err
	}

	return prd, nil
}
