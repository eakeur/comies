package stocking

import (
	"comies/app/sdk/throw"
	"comies/app/sdk/types"
	"context"
)

func (w workflow) RemoveStock(ctx context.Context, id types.ID) error {

	err := w.stocks.Remove(ctx, id)
	if err != nil {
		return throw.Error(err)
	}

	return nil
}
