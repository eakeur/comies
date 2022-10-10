package ordering

import (
	"comies/app/core/id"
	"comies/app/core/order"
	"comies/app/data/orders"
	"context"
	"time"
)

func SetOrderStatus(ctx context.Context, id id.ID, st order.Status) error {
	_, err := orders.UpdateFlow(ctx, order.FlowUpdate{
		OrderID:    id,
		Status:     st,
		OccurredAt: time.Now().UTC(),
	})
	if err != nil {
		return err
	}

	return nil
}
