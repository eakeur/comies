package product

import (
	"context"
	"gomies/app/sdk/fault"
	"gomies/app/sdk/types"
)

func (w workflow) UpdateReservation(ctx context.Context, reservationID types.ID, consume bool) error {
	const operation = "Workflows.Product.UpdateReservation"

	if consume {
		err := w.stocks.ConsumeResources(ctx, reservationID)
		if err != nil {
			return fault.Wrap(err, operation)
		}
	}

	err := w.stocks.FreeResources(ctx, reservationID)
	if err != nil {
		return fault.Wrap(err, operation)
	}

	return nil
}
