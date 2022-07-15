package stocking

import (
	"comies/app/sdk/fault"
	"comies/app/sdk/types"
	"context"
)

func (w workflow) RemoveStock(ctx context.Context, id types.ID) error {

	err := w.stocks.Remove(ctx, id)
	if err != nil {
		return fault.Wrap(err)
	}

	return nil
}
