package menu

import (
	"comies/app/core/entities/movement"
	"comies/app/core/types"
	"context"
)

func (w workflow) GetMovementsBalance(ctx context.Context, filter movement.Filter) (types.Quantity, error) {
	if err := filter.Validate(); err != nil {
		return 0, err
	}

	actual, err := w.movements.GetBalanceByProductID(ctx, filter.ProductID, filter)
	if err != nil {
		return 0, err
	}

	return actual, nil
}
