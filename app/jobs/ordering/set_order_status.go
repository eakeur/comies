package ordering

import (
	"comies/app/core/entities/order"
	"comies/app/core/types"
	"context"
	"time"
)

func (w jobs) SetOrderStatus(ctx context.Context, id types.ID, st order.Status) error {
	_, err := w.orders.UpdateFlow(ctx, order.FlowUpdate{
		OrderID:    id,
		Status:     st,
		OccurredAt: time.Now().UTC(),
	})
	if err != nil {
		return err
	}

	return nil
}
