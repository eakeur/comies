package menu

import (
	"context"
	product2 "gomies/app/core/entities/product"
	"gomies/app/sdk/fault"
)

func (w workflow) GetProduct(ctx context.Context, ext product2.Key) (product2.Product, error) {

	prod, err := w.products.GetByID(ctx, ext.ID)
	if err != nil {
		return product2.Product{}, fault.Wrap(err)
	}
	return prod, nil
}
