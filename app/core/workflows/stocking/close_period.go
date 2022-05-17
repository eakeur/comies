package stocking

import (
	"context"
	"gomies/app/core/entities/movement"
	"gomies/app/sdk/fault"
)

func (w workflow) ClosePeriod(ctx context.Context, filter movement.Filter) error {

	if err := filter.Validate(); err != nil {
		return fault.Wrap(err)
	}

	err := w.movements.Archive(ctx, filter)
	if err != nil {
		return fault.Wrap(err)
	}

	return nil
}
