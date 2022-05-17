package menu

import (
	"context"
	"gomies/app/sdk/fault"
	"gomies/app/sdk/types"
)

func (w workflow) UpdateReservation(ctx context.Context, reservationID types.ID, consume bool) error {

	if consume {
		err := w.stocks.ConsumeResources(ctx, reservationID)
		if err != nil {
			return fault.Wrap(err)
		}
	}

	err := w.stocks.FreeResources(ctx, reservationID)
	if err != nil {
		return fault.Wrap(err)
	}

	return nil
}
