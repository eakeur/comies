package menu

import (
	"comies/app/core/movement"
	"comies/app/core/types"
	"comies/app/data/movements"
	"context"
)

func GetProductBalance(ctx context.Context, filter movement.Filter) (types.Quantity, error) {
	if err := movement.ValidateFilter(filter); err != nil {
		return 0, err
	}

	return movements.GetBalance(ctx, filter)
}
