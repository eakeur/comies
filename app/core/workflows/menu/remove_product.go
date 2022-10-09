package menu

import (
	"comies/app/core/types"
	"context"
)

func (w workflow) RemoveProduct(ctx context.Context, id types.ID) error {

	err := w.products.Remove(ctx, id)
	if err != nil {
		return err
	}

	return err

}
