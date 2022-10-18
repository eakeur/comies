package menu

import (
	"comies/app/core/menu"
	"comies/app/data/movements"
	"context"
)

func ListMovements(ctx context.Context, filter menu.MovementFilter) ([]menu.Movement, error) {
	if err := menu.ValidateMovementFilter(filter); err != nil {
		return nil, err
	}

	return movements.List(ctx, filter)
}
