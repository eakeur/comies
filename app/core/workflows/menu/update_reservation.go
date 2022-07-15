package menu

import (
	"comies/app/sdk/fault"
	"comies/app/sdk/types"
	"context"
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
