package menu

import (
	"comies/app/core/entities/product"
	"comies/app/sdk/throw"
	"context"
)

func (w workflow) GetProductByCode(ctx context.Context, code string) (product.Product, error) {

	prod, err := w.products.GetByCode(ctx, code)
	if err != nil {
		return product.Product{}, throw.Error(err)
	}
	return prod, nil
}
