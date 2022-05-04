package stock

import (
	"context"
	"gomies/app/sdk/fault"
	"gomies/app/sdk/types"
)

func (w workflow) ConsumeResources(ctx context.Context, reservationID types.ID) error {
	const operation = "Workflows.Stock.ConsumeResources"

	err := w.stocks.UpdateReserved(ctx, reservationID)
	if err != nil {
		return fault.Wrap(err, operation)
	}

	return nil
}
