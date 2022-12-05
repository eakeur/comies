package menu

import (
	"comies/core/menu/movement"
	"comies/core/types"
	"context"
)

func (w jobs) GetProductStockBalance(ctx context.Context, filter movement.Filter) (types.Quantity, error) {
	err := filter.Validate()
	if err != nil {
		return 0, err
	}

	return w.movements.Balance(ctx, filter)
}
