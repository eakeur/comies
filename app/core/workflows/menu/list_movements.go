package menu

import (
	"comies/app/core/movement"
	"comies/app/data/movements"
	"context"
)

func ListMovements(ctx context.Context, filter movement.Filter) ([]movement.Movement, error) {
	if err := filter.Validate(); err != nil {
		return nil, err
	}

	movements, err := movements.ListByProductID(ctx, filter.ProductID, filter)
	if err != nil {
		return nil, err
	}

	return movements, nil
}
