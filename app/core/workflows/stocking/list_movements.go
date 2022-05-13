package stocking

import (
	"context"
	"gomies/app/core/entities/movement"
	"gomies/app/sdk/fault"
)

func (w workflow) ListMovements(ctx context.Context, filter movement.Filter) ([]movement.Movement, int, error) {
	const operation = "Workflows.Stock.ListMovements"

	if err := filter.Validate(); err != nil {
		return []movement.Movement{}, 0, fault.Wrap(err, operation)
	}

	movements, count, err := w.movements.List(ctx, filter)
	if err != nil {
		return []movement.Movement{}, 0, fault.Wrap(err, operation)
	}

	return movements, count, nil
}
