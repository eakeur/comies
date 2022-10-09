package menu

import (
	"comies/app/core/entities/product"
	"context"
)

func (w workflow) UpdateProduct(ctx context.Context, prd product.Product) error {

	if err := prd.Validate(); err != nil {
		return err
	}

	err := w.products.Update(ctx, prd)
	if err != nil {
		return err
	}

	return nil
}
