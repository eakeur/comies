package menu

import (
	"comies/app/core/types"
	"context"
)

func (w jobs) RemoveProduct(ctx context.Context, id types.ID) error {
	if err := id.Validate(); err != nil {
		return err
	}
	
	return w.products.Remove(ctx, id)
}
