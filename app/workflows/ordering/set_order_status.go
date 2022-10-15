package ordering

import (
	"comies/app/core/id"
	"comies/app/core/ordering"
	"comies/app/data/orders"
	"context"
)

func SetOrderStatus(ctx context.Context, id id.ID, st ordering.Status) error {
	if err := ordering.ValidateOrderStatus(st); err != nil {
		return err
	}
	
	return orders.UpdateFlow(ctx, ordering.NewFlow(id, st))
}
