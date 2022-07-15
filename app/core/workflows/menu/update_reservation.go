package menu

import (
	"comies/app/sdk/throw"
	"comies/app/sdk/types"
	"context"
)

func (w workflow) UpdateReservation(ctx context.Context, reservationID types.ID, consume bool) error {

	if consume {
		err := w.stocks.ConsumeResources(ctx, reservationID)
		if err != nil {
			return throw.Error(err)
		}
	}

	err := w.stocks.FreeResources(ctx, reservationID)
	if err != nil {
		return throw.Error(err)
	}

	return nil
}
