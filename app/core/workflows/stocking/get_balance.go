package stocking

import (
	"comies/app/core/entities/movement"
	"comies/app/sdk/throw"
	"comies/app/sdk/types"
	"context"
)

func (w workflow) GetBalance(ctx context.Context, filter movement.Filter) (types.Quantity, error) {

	if err := filter.Validate(); err != nil {
		return 0, throw.Error(err)
	}

	actual, err := w.movements.GetBalanceByResourceID(ctx, filter.ResourceID, filter)
	if err != nil {
		return 0, throw.Error(err)
	}

	return actual, nil

}
