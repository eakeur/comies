package menu

import (
	"comies/app/core/entities/product"
	"context"
)

func (w workflow) GetProductByCode(ctx context.Context, code string) (product.Product, error) {

	prod, err := w.products.GetByCode(ctx, code)
	if err != nil {
		return product.Product{}, err
	}
	return prod, nil
}
