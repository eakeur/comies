package menu

import (
	"comies/app/core/entities/product"
	"comies/app/sdk/throw"
	"context"
)

func (w workflow) UpdateProduct(ctx context.Context, prd product.Product) error {

	if err := prd.Validate(); err != nil {
		return throw.Error(err)
	}

	err := w.products.Update(ctx, prd)
	if err != nil {
		return throw.Error(err)
	}

	return nil
}
