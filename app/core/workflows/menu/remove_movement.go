package menu

import (
	"comies/app/core/throw"
	"comies/app/core/types"
	"context"
)

func (w workflow) RemoveMovement(ctx context.Context, movementID types.ID) error {
	if movementID.Empty() {
		return throw.ErrMissingID
	}

	err := w.movements.Remove(ctx, movementID)
	if err != nil {
		return err
	}

	return nil
}
