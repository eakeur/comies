package stocking

import (
	"context"
	"gomies/app/sdk/fault"
	"gomies/app/sdk/types"
)

func (w workflow) RemoveMovement(ctx context.Context, resourceID types.ID, movementID types.ID) error {

	if resourceID.Empty() || movementID.Empty() {
		return fault.Wrap(fault.ErrMissingID)
	}

	err := w.movements.Remove(ctx, movementID)
	if err != nil {
		return fault.Wrap(err)
	}

	return nil
}
