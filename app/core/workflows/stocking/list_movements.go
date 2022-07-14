package stocking

import (
	"context"
	"gomies/app/core/entities/movement"
	"gomies/app/sdk/fault"
)

func (w workflow) ListMovements(ctx context.Context, filter movement.Filter) ([]movement.Movement, error) {

	if err := filter.Validate(); err != nil {
		return []movement.Movement{}, fault.Wrap(err)
	}

	movements, err := w.movements.ListByResourceID(ctx, filter.ResourceID, filter)
	if err != nil {
		return []movement.Movement{}, fault.Wrap(err)
	}

	return movements, nil
}
