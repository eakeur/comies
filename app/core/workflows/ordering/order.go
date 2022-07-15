package ordering

import (
	"comies/app/core/entities/order"
	"comies/app/sdk/throw"
	"context"
	"time"
)

func (w workflow) Order(ctx context.Context, o OrderConfirmation) (order.Order, error) {

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

	for _, item := range items {
		err := w.products.UpdateResources(ctx, item.ID, true)
		if err != nil {
			return order.Order{}, throw.Error(err).Params(map[string]interface{}{
				"item_id": item.ID,
				"consume": true,
			})
		}
	}

	return ord, nil

}
