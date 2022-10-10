package menu

import (
	"comies/app/core/id"
	"comies/app/data/movements"
	"context"
)

func UpdateReservation(ctx context.Context, reservationID id.ID, consume bool) error {

	if consume {
		err := movements.SetOutputType(ctx, reservationID)
		if err != nil {
			return err
		}

	}

	err := movements.RemoveReserved(ctx, reservationID)
	if err != nil {
		return err
	}

	return nil
}
