package menu

import (
	"comies/app/sdk/throw"
	"comies/app/sdk/types"
	"context"
)

func (w workflow) UpdateReservation(ctx context.Context, reservationID types.ID, consume bool) error {

	if consume {
		err := w.movements.SetOutputType(ctx, reservationID)
		if err != nil {
			return throw.Error(err)
		}

	}

	err := w.movements.RemoveReserved(ctx, reservationID)
	if err != nil {
		return throw.Error(err)
	}

	return nil
}
