package ordering

import (
	"comies/app/core/entities/order"
	"context"
)

func (w workflow) Channel(ctx context.Context) (chan OrderNotification, error) {
	ch, ok := w.channel[0]
	if ok {
		return ch, nil
	}

	ch, err := w.createChannel(ctx)
	if err != nil {
		return nil, err
	}

	return ch, nil
}

func (w workflow) createChannel(ctx context.Context) (chan OrderNotification, error) {
	orders, err := w.orders.List(ctx, order.Filter{
		Status: []order.Status{
			order.PendingStatus,
			order.PreparingStatus,
			order.WaitingTakeoutStatus,
			order.WaitingDeliveryStatus,
			order.DeliveringStatus,
		},
	})
	if err != nil {
		return nil, err
	}

	ch := make(chan OrderNotification)
	for _, o := range orders {
		items, err := w.items.List(ctx, o.ID)
		if err != nil {
			return nil, err
		}

		ch <- OrderNotification{
			Order: o,
			Items: items,
		}
	}

	w.channel[0] = ch
	return ch, nil
}

func (w workflow) sendToChannel(not OrderNotification) {
	ch, ok := w.channel[0]
	if !ok {
		return
	}

	ch <- not
}
