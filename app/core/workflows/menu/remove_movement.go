package menu

import (
	"comies/app/sdk/throw"
	"comies/app/sdk/types"
	"context"
)

func (w workflow) RemoveMovement(ctx context.Context, movementID types.ID) error {
	if movementID.Empty() {
		return throw.Error(throw.ErrMissingID)
	}

	err := w.movements.Remove(ctx, movementID)
	if err != nil {
		return throw.Error(err)
	}

	return nil
}
