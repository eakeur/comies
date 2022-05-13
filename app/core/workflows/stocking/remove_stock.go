package stocking

import (
	"context"
	"gomies/app/sdk/fault"
	"gomies/app/sdk/types"
)

func (w workflow) RemoveStock(ctx context.Context, id types.ID) error {
	const operation = "Workflows.Stock.RemoveStock"

	err := w.stocks.RemoveStock(ctx, id)
	if err != nil {
		return fault.Wrap(err, operation)
	}

	return nil
}
