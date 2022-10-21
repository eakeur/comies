package ordering

import (
	"comies/app/core/ordering"
	"comies/app/core/types"
	"comies/app/data/items"
	"comies/app/data/orders"
	"context"
)

func RemoveItemFromPendingOrder(ctx context.Context, orderID, itemID types.ID) error {
	if err := types.ValidateID(itemID); err != nil {
		return err
	}

	status, err := orders.GetStatus(ctx, orderID)
	if err != nil {
		return err
	}

	if status > ordering.PendingOrderStatus {
		return nil
	}

	return items.Remove(ctx, itemID)
}
