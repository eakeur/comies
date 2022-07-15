package menu

import (
	"context"
	"gomies/app/core/entities/product"
	"gomies/app/sdk/fault"
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
