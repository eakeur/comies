package stocking

import (
	"context"
	"gomies/app/sdk/fault"
	"gomies/app/sdk/types"
)

func (w workflow) FreeResources(ctx context.Context, reservationID types.ID) error {

	err := w.movements.RemoveReserved(ctx, reservationID)
	if err != nil {
		return fault.Wrap(err)
	}

	return nil
}
