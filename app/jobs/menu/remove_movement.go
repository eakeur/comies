package menu

import (
	"comies/app/core/types"
	"comies/app/data/movements"
	"context"
)

func RemoveMovement(ctx context.Context, movementID types.ID) error {
	if err := types.ValidateID(movementID); err != nil {
		return err
	}

	err := movements.Remove(ctx, movementID)
	if err != nil {
		return err
	}

	return nil
}
