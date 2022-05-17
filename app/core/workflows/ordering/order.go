package ordering

import (
	"context"
	"gomies/app/core/entities/order"
	"gomies/app/sdk/fault"
	"sync"
)

func (w workflow) Order(ctx context.Context, o OrderConfirmation) (order.Order, error) {

	ord, err := w.orders.GetOrder(ctx, o.OrderID)
	if err != nil {
		return order.Order{}, fault.Wrap(err).Params(map[string]interface{}{
			"order_id": o.OrderID,
		})
	}

	if ord.Status >= order.PreparingStatus {
		return order.Order{}, fault.Wrap(order.ErrOrderAlreadyOrdered).Params(map[string]interface{}{
			"order_status": ord.Status,
		})
	}

	items, err := w.orders.ListItems(ctx, o.OrderID)
	if err != nil {
		return order.Order{}, fault.Wrap(err).Params(map[string]interface{}{
			"order_id": o.OrderID,
		})
	}

	if len(items) <= 0 {
		return order.Order{}, fault.Wrap(order.ErrInvalidNumberOfItems).Params(map[string]interface{}{
			"order_id": o.OrderID,
		})
	}

	failures := make(chan error)
	wg := sync.WaitGroup{}

	wg.Add(3)
	go func() {
		defer wg.Done()
		failures <- w.orders.UpdateOrderAddressID(ctx, o.OrderID, o.AddressID)
	}()

	go func() {
		defer wg.Done()
		failures <- w.orders.UpdateOrderDeliveryMode(ctx, o.OrderID, o.DeliveryMode)
	}()

	go func() {
		defer wg.Done()
		failures <- w.orders.UpdateOrderStatus(ctx, o.OrderID, order.PreparingStatus)
	}()
	wg.Wait()
	close(failures)

	if err := <-failures; err != nil {
		return order.Order{}, fault.Wrap(err).Params(map[string]interface{}{
			"order_id":      o.OrderID,
			"address_id":    o.AddressID,
			"delivery_mode": o.DeliveryMode,
			"order_status":  order.PreparingStatus,
		})
	}

	for _, item := range items {
		err := w.products.UpdateResources(ctx, item.ID, true)
		if err != nil {
			return order.Order{}, fault.Wrap(err).Params(map[string]interface{}{
				"order_id": item.ID,
				"consume":  true,
			})
		}
	}

	return ord, nil

}
