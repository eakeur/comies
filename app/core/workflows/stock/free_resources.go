package stock

import (
	"context"
	"gomies/app/sdk/fault"
	"gomies/app/sdk/types"
)

func (w workflow) FreeResources(ctx context.Context, reservationID types.ID) error {
	const operation = "Workflows.Stock.FreeResources"

	err := w.stocks.RemoveReserved(ctx, reservationID)
	if err != nil {
		return fault.Wrap(err, operation)
	}

	return nil
}
