package menu

import (
	"comies/app/core/entities/product"
	"comies/app/sdk/fault"
	"context"
)

func (w workflow) UpdateProduct(ctx context.Context, prd product.Product) error {

	if err := prd.Validate(); err != nil {
		return fault.Wrap(err)
	}

	err := w.products.Update(ctx, prd)
	if err != nil {
		return fault.Wrap(err)
	}

	return nil
}
