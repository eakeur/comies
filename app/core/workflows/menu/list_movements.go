package menu

import (
	"comies/app/core/entities/movement"
	"comies/app/sdk/throw"
	"context"
)

func (w workflow) ListMovements(ctx context.Context, filter movement.Filter) ([]movement.Movement, error) {
	if err := filter.Validate(); err != nil {
		return nil, throw.Error(err)
	}

	movements, err := w.movements.ListByProductID(ctx, filter.ProductID, filter)
	if err != nil {
		return nil, throw.Error(err)
	}

	return movements, nil
}
