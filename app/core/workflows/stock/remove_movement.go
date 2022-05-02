package stock

import (
	"context"
	"gomies/app/core/entities/stock"
	"gomies/app/sdk/fault"
	"gomies/app/sdk/types"
)

func (w workflow) RemoveMovement(ctx context.Context, resourceID types.ID, movementID types.ID) error {
	const operation = "Workflows.Stock.RemoveMovement"

	if resourceID.Empty() || movementID.Empty() {
		return fault.Wrap(stock.ErrMissingResourceID, operation)
	}

	err := w.stocks.RemoveMovement(ctx, resourceID, movementID)
	if err != nil {
		return fault.Wrap(err, operation)
	}

	return nil
}
