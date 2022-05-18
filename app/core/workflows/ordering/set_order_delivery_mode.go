package ordering

import (
	"context"
	"gomies/app/core/entities/order"
	"gomies/app/sdk/fault"
	"gomies/app/sdk/types"
)

func (w workflow) SetOrderDeliveryMode(ctx context.Context, id types.ID, deliveryMode order.DeliveryMode) error {
	err := w.orders.SetDeliveryMode(ctx, id, deliveryMode)
	if err != nil {
		return fault.Wrap(err)
	}

	return nil
}
