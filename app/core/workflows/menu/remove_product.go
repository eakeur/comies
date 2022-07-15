package menu

import (
	"comies/app/core/entities/product"
	"comies/app/sdk/throw"
	"context"
)

func (w workflow) RemoveProduct(ctx context.Context, ext product.Key) error {

	err := w.products.Remove(ctx, ext.ID)
	if err != nil {
		return throw.Error(err)
	}

	return err

}
