package menu

import (
	"comies/app/core/id"
	"comies/app/data/movements"
	"context"
)

func RemoveMovement(ctx context.Context, movementID id.ID) error {
	if err := id.ValidateID(movementID); err != nil {
		return err
	}

	err := movements.Remove(ctx, movementID)
	if err != nil {
		return err
	}

	return nil
}
