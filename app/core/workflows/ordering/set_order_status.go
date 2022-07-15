package ordering

import (
	"comies/app/core/entities/order"
	"comies/app/sdk/throw"
	"comies/app/sdk/types"
	"context"
	"time"
)

func (w workflow) SetOrderStatus(ctx context.Context, id types.ID, st order.Status) error {
	_, err := w.orders.UpdateFlow(ctx, order.FlowUpdate{
		OrderID:    id,
		Status:     st,
		OccurredAt: time.Now().UTC(),
	})
	if err != nil {
		return throw.Error(err)
	}

	return nil
}
