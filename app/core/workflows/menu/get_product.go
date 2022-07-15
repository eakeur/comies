package menu

import (
	product2 "comies/app/core/entities/product"
	"comies/app/sdk/throw"
	"context"
)

func (w workflow) GetProduct(ctx context.Context, ext product2.Key) (product2.Product, error) {

	prod, err := w.products.GetByID(ctx, ext.ID)
	if err != nil {
		return product2.Product{}, throw.Error(err)
	}
	return prod, nil
}
