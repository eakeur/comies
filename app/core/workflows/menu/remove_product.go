package menu

import (
	"comies/app/core/throw"
	"comies/app/core/types"
	"context"
)

func (w workflow) RemoveProduct(ctx context.Context, id types.ID) error {

	err := w.products.Remove(ctx, id)
	if err != nil {
		return throw.Error(err)
	}

	return err

}
