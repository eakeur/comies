package ordering

import (
	"comies/app/core/ordering"
	"comies/app/data/ids"
	"comies/app/data/orders"
	"context"
)

func InitializeOrder(ctx context.Context) (ordering.Order, error) {
	o := ordering.Order{
		ID:     ids.Create(),
		Status: ordering.InTheCartOrderStatus,
	}
	return o, orders.Create(ctx, o)
}
