package ordering

import (
	"comies/app/core/order"
	"comies/app/data/orders"
	"context"
)

func ListOrders(ctx context.Context, f order.Filter) ([]order.Order, error) {
	list, err := orders.List(ctx, f)
	if err != nil {
		return nil, err
	}

	return list, nil
}
