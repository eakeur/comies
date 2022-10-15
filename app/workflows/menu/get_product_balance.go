package menu

import (
	"comies/app/core/menu"
	"comies/app/core/types"
	"comies/app/data/movements"
	"context"
)

func GetProductBalance(ctx context.Context, filter menu.MovementFilter) (types.Quantity, error) {
	if err := menu.ValidateMovementFilter(filter); err != nil {
		return 0, err
	}

	return movements.GetBalance(ctx, filter)
}
