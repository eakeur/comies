package menu

import (
	"comies/app/core/id"
	"comies/app/core/types"
	"comies/app/data/movements"
	"context"
)

func RemoveMovement(ctx context.Context, movementID id.ID) error {
	if movementID.Empty() {
		return types.ErrMissingID
	}

	err := movements.Remove(ctx, movementID)
	if err != nil {
		return err
	}

	return nil
}
