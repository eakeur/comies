package ordering

import (
	"comies/app/core/ordering"
	"comies/app/core/types"
	"comies/app/data/orders"
	"context"
)

func SetOrderDeliveryMode(ctx context.Context, id types.ID, deliveryType ordering.Type) error {
	if err := ordering.ValidateDeliveryType(deliveryType); err != nil {
		return err
	}

	return orders.SetDeliveryType(ctx, id, deliveryType)
}
