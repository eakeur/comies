package stocking

import (
	"context"
	"gomies/app/sdk/fault"
	"gomies/app/sdk/types"
)

func (w workflow) RemoveMovement(ctx context.Context, resourceID types.ID, movementID types.ID) error {
	const operation = "Workflows.Stock.RemoveMovement"

	if resourceID.Empty() || movementID.Empty() {
		return fault.Wrap(fault.ErrMissingUID, operation)
	}

	err := w.movements.Remove(ctx, resourceID, movementID)
	if err != nil {
		return fault.Wrap(err, operation)
	}

	return nil
}
