package stocking

import (
	"context"
	"gomies/app/sdk/fault"
	"gomies/app/sdk/types"
)

func (w workflow) ConsumeResources(ctx context.Context, reservationID types.ID) error {

	err := w.movements.SetOutputType(ctx, reservationID)
	if err != nil {
		return fault.Wrap(err)
	}

	return nil
}
