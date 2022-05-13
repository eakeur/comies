package stocking

import (
	"context"
	"gomies/app/core/entities/movement"
	"gomies/app/sdk/fault"
)

func (w workflow) ClosePeriod(ctx context.Context, filter movement.Filter) error {
	const operation = "Workflows.Stock.ClosePeriod"

	if err := filter.Validate(); err != nil {
		return fault.Wrap(err, operation)
	}

	err := w.movements.Archive(ctx, filter)
	if err != nil {
		return fault.Wrap(err, operation)
	}

	return nil
}
