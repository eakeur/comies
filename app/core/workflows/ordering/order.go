package ordering

import (
	"comies/app/core/entities/order"
	"comies/app/core/throw"
	"context"
	"time"
)

func (w workflow) Order(ctx context.Context, o OrderConfirmation) (order.Order, error) {

	var consume bool

	ord, err := w.orders.GetByID(ctx, o.OrderID)
	if err != nil {
		return order.Order{}, throw.Error(err).Params(map[string]interface{}{
			"order_id": o.OrderID,
		})
	}

	if ord.Status >= order.PreparingStatus {
		return order.Order{}, throw.Error(order.ErrAlreadyOrdered).Params(map[string]interface{}{
			"order_status": ord.Status,
		})
	}

	items, err := w.items.List(ctx, o.OrderID)
	if err != nil {
		return order.Order{}, throw.Error(err).Params(map[string]interface{}{
			"order_id": o.OrderID,
		})
	}

	if len(items) <= 0 {
		return order.Order{}, throw.Error(order.ErrInvalidNumberOfItems).Params(map[string]interface{}{
			"order_id": o.OrderID,
		})
	}

	defer func() {
		go func() {
			w.sendToChannel(OrderNotification{
				Order: ord,
				Items: items,
			})
			for _, item := range items {
				err := w.products.UpdateReservation(ctx, item.ID, consume)
				if err != nil {
					err = throw.Error(err).Params(map[string]interface{}{
						"item_id": item.ID,
						"consume": consume,
					})
				}
			}
		}()
	}()

	if err := w.orders.SetDeliveryMode(ctx, o.OrderID, o.DeliveryMode); err != nil {
		return order.Order{}, throw.Error(err).Params(map[string]interface{}{
			"order_id":      o.OrderID,
			"delivery_mode": o.DeliveryMode,
		})
	}

	if _, err = w.orders.UpdateFlow(ctx, order.FlowUpdate{
		OrderID:    o.OrderID,
		Status:     order.PreparingStatus,
		OccurredAt: time.Now().UTC(),
	}); err != nil {
		return order.Order{}, throw.Error(err).Params(map[string]interface{}{
			"order_id":     o.OrderID,
			"order_status": order.PreparingStatus,
		})
	}

	consume = true

	return ord, nil

}
