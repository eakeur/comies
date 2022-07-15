package stocking

import (
	"comies/app/core/entities/movement"
	"comies/app/sdk/throw"
	"context"
)

func (w workflow) ListMovements(ctx context.Context, filter movement.Filter) ([]movement.Movement, error) {

	if err := filter.Validate(); err != nil {
		return []movement.Movement{}, throw.Error(err)
	}

	movements, err := w.movements.ListByResourceID(ctx, filter.ResourceID, filter)
	if err != nil {
		return []movement.Movement{}, throw.Error(err)
	}

	return movements, nil
}
