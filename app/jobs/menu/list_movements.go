package menu

import (
	"comies/app/core/movement"
	"comies/app/data/movements"
	"context"
)

func ListMovements(ctx context.Context, filter movement.Filter) ([]movement.Movement, error) {
	if err := movement.ValidateFilter(filter); err != nil {
		return nil, err
	}

	return movements.List(ctx, filter)
}
