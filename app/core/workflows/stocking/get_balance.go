package stocking

import (
	"context"
	"gomies/app/core/entities/movement"
	"gomies/app/sdk/fault"
	"gomies/app/sdk/types"
)

func (w workflow) GetBalance(ctx context.Context, filter movement.Filter) (types.Quantity, error) {

	if err := filter.Validate(); err != nil {
		return 0, fault.Wrap(err)
	}

	actual, err := w.movements.GetBalance(ctx, filter)
	if err != nil {
		return 0, fault.Wrap(err)
	}

	return actual, nil

}
