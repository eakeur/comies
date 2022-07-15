package menu

import (
	"context"
	"gomies/app/core/entities/product"
	"gomies/app/sdk/fault"
)

func (w workflow) RemoveProduct(ctx context.Context, ext product.Key) error {

	err := w.products.Remove(ctx, ext.ID)
	if err != nil {
		return fault.Wrap(err)
	}

	return err

}
