package menu

import (
	"comies/app/core/entities/product"
	"comies/app/sdk/throw"
	"comies/app/sdk/types"
	"context"
)

func (w workflow) GetProductByID(ctx context.Context, id types.ID) (product.Product, error) {

	prod, err := w.products.GetByID(ctx, id)
	if err != nil {
		return product.Product{}, throw.Error(err)
	}
	return prod, nil
}
