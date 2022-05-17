package stocking

import (
	"context"
	"gomies/app/sdk/fault"
	"gomies/app/sdk/types"
)

func (w workflow) RemoveStock(ctx context.Context, id types.ID) error {

	err := w.stocks.RemoveStock(ctx, id)
	if err != nil {
		return fault.Wrap(err)
	}

	return nil
}
