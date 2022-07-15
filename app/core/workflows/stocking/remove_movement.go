package stocking

import (
	"comies/app/sdk/throw"
	"comies/app/sdk/types"
	"context"
)

func (w workflow) RemoveMovement(ctx context.Context, resourceID types.ID, movementID types.ID) error {

	if resourceID.Empty() || movementID.Empty() {
		return throw.Error(throw.ErrMissingID)
	}

	err := w.movements.Remove(ctx, movementID)
	if err != nil {
		return throw.Error(err)
	}

	return nil
}
