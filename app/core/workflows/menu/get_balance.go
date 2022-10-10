package menu

import (
	"comies/app/core/movement"
	"comies/app/core/types"
	"comies/app/data/movements"
	"context"
)

func GetMovementsBalance(ctx context.Context, filter movement.Filter) (types.Quantity, error) {
	if err := filter.Validate(); err != nil {
		return 0, err
	}

	actual, err := movements.GetBalanceByProductID(ctx, filter.ProductID, filter)
	if err != nil {
		return 0, err
	}

	return actual, nil
}
