package menu

import (
	"comies/app/core/types"
	"context"
)

func (w workflow) UpdateReservation(ctx context.Context, reservationID types.ID, consume bool) error {

	if consume {
		err := w.movements.SetOutputType(ctx, reservationID)
		if err != nil {
			return err
		}

	}

	err := w.movements.RemoveReserved(ctx, reservationID)
	if err != nil {
		return err
	}

	return nil
}
