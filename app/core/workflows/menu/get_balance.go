package menu

import (
	"comies/app/core/entities/movement"
	"comies/app/core/throw"
	"comies/app/core/types"
	"context"
)

func (w workflow) GetMovementsBalance(ctx context.Context, filter movement.Filter) (types.Quantity, error) {
	if err := filter.Validate(); err != nil {
		return 0, throw.Error(err)
	}

	actual, err := w.movements.GetBalanceByProductID(ctx, filter.ProductID, filter)
	if err != nil {
		return 0, throw.Error(err)
	}

	return actual, nil
}
