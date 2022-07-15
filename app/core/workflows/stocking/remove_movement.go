package stocking

import (
	"comies/app/sdk/fault"
	"comies/app/sdk/types"
	"context"
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
