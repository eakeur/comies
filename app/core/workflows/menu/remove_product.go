package menu

import (
	"comies/app/core/entities/product"
	"comies/app/sdk/fault"
	"context"
)

func (w workflow) RemoveProduct(ctx context.Context, ext product.Key) error {

	err := w.products.Remove(ctx, ext.ID)
	if err != nil {
		return fault.Wrap(err)
	}

	return err

}
