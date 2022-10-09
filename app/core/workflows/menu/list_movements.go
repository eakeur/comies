package menu

import (
	"comies/app/core/entities/movement"
	"context"
)

func (w workflow) ListMovements(ctx context.Context, filter movement.Filter) ([]movement.Movement, error) {
	if err := filter.Validate(); err != nil {
		return nil, err
	}

	movements, err := w.movements.ListByProductID(ctx, filter.ProductID, filter)
	if err != nil {
		return nil, err
	}

	return movements, nil
}
