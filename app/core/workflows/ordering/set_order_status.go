package ordering

import (
	"context"
	"gomies/app/core/entities/order"
	"gomies/app/sdk/fault"
	"gomies/app/sdk/types"
	"time"
)

func (w workflow) SetOrderStatus(ctx context.Context, id types.ID, st order.Status) error {
	_, err := w.orders.UpdateFlow(ctx, order.FlowUpdate{
		OrderID:    id,
		Status:     st,
		OccurredAt: time.Now().UTC(),
	})
	if err != nil {
		return fault.Wrap(err)
	}

	return nil
}
