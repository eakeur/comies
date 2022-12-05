package menu

import (
	"comies/core/types"
	"context"
)

func (w jobs) RemoveMovement(ctx context.Context, id types.ID) error {

	if err := id.Validate(); err != nil {
		return err
	}

	return w.movements.Remove(ctx, id)
}
