package stocking

import (
	"comies/app/sdk/fault"
	"comies/app/sdk/types"
	"context"
)

func (w workflow) FreeResources(ctx context.Context, reservationID types.ID) error {

	err := w.movements.RemoveReserved(ctx, reservationID)
	if err != nil {
		return fault.Wrap(err)
	}

	return nil
}
