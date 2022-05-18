package ordering

import (
	"context"
	"gomies/app/core/entities/order"
	"gomies/app/sdk/fault"
	"gomies/app/sdk/types"
)

func (w workflow) SetOrderStatus(ctx context.Context, id types.ID, st order.Status) error {
	err := w.orders.SetStatus(ctx, id, st)
	if err != nil {
		return fault.Wrap(err)
	}

	return nil
}
