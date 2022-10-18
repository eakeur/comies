package menu

import (
	"comies/app/core/id"
	"comies/app/data/movements"
	"context"
)

func ConfirmReservation(ctx context.Context, reservationID id.ID) error {
	return movements.SetOutputType(ctx, reservationID)
}

func CancelReservation(ctx context.Context, reservationID id.ID) error {
	return movements.RemoveReserved(ctx, reservationID)
}
