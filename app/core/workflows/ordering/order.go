package ordering

import (
	"comies/app/core/order"
	"comies/app/core/workflows/menu"
	"comies/app/data/items"
	"comies/app/data/orders"
	"context"
	"time"
)

func Order(ctx context.Context, o OrderConfirmation) (order.Order, error) {

	var consume bool

	ord, err := orders.GetByID(ctx, o.OrderID)
	if err != nil {
		return order.Order{}, err
	}

	if ord.Status >= order.PreparingStatus {
		return order.Order{}, order.ErrAlreadyOrdered
	}

	items, err := items.List(ctx, o.OrderID)
	if err != nil {
		return order.Order{}, err
	}

	if len(items) <= 0 {
		return order.Order{}, order.ErrInvalidNumberOfItems
	}

	defer func() {
		go func() {
			sendToChannel(OrderNotification{
				Order: ord,
				Items: items,
			})
			for _, item := range items {
				menu.UpdateReservation(ctx, item.ID, consume)
			}
		}()
	}()

	if err := orders.SetDeliveryMode(ctx, o.OrderID, o.DeliveryMode); err != nil {
		return order.Order{}, err
	}

	if _, err = orders.UpdateFlow(ctx, order.FlowUpdate{
		OrderID:    o.OrderID,
		Status:     order.PreparingStatus,
		OccurredAt: time.Now().UTC(),
	}); err != nil {
		return order.Order{}, err
	}

	consume = true

	return ord, nil

}
