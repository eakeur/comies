package ordering

import (
	"comies/app/core/entities/order"
	"comies/app/sdk/fault"
	"comies/app/sdk/types"
	"context"
)

func (w workflow) SetOrderDeliveryMode(ctx context.Context, id types.ID, deliveryMode order.DeliveryMode) error {
	err := w.orders.SetDeliveryMode(ctx, id, deliveryMode)
	if err != nil {
		return fault.Wrap(err)
	}

	return nil
}
