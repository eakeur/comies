package ordering

import (
	"comies/app/core/ordering"
	"comies/app/data/ids"
	"comies/app/data/orders"
	"context"
)

func InitializeOrder(ctx context.Context) (ordering.Order, error) {
	return orders.Create(ctx, ordering.Order{
		ID: ids.Create(),
		Status: ordering.InTheCartOrderStatus,
	})
}
