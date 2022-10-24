package ordering

import (
	"comies/app/core/entities/order"
	"context"
	"time"
)

func (w jobs) Order(ctx context.Context, o OrderConfirmation) (order.Order, error) {

	var consume bool

	ord, err := w.orders.GetByID(ctx, o.OrderID)
	if err != nil {
		return order.Order{}, err
	}

	if ord.Status >= order.PreparingStatus {
		return order.Order{}, order.ErrAlreadyOrdered
	}

	items, err := w.items.List(ctx, o.OrderID)
	if err != nil {
		return order.Order{}, err
	}

	if len(items) <= 0 {
		return order.Order{}, order.ErrInvalidNumberOfItems
	}

	defer func() {
		go func() {
			w.sendToChannel(OrderNotification{
				Order: ord,
				Items: items,
			})
			for _, item := range items {
				w.products.UpdateReservation(ctx, item.ID, consume)
			}
		}()
	}()

	if err := w.orders.SetDeliveryMode(ctx, o.OrderID, o.DeliveryMode); err != nil {
		return order.Order{}, err
	}

	if _, err = w.orders.UpdateFlow(ctx, order.FlowUpdate{
		OrderID:    o.OrderID,
		Status:     order.PreparingStatus,
		OccurredAt: time.Now().UTC(),
	}); err != nil {
		return order.Order{}, err
	}

	consume = true

	return ord, nil

}
