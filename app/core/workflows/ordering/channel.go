package ordering

import (
	"comies/app/core/id"
	"comies/app/core/order"
	"comies/app/data/items"
	"comies/app/data/orders"
	"context"
)

var channel = make(map[id.ID]chan OrderNotification)

func Channel(ctx context.Context) (chan OrderNotification, error) {
	ch, ok := channel[0]
	if ok {
		return ch, nil
	}

	ch, err := createChannel(ctx)
	if err != nil {
		return nil, err
	}

	return ch, nil
}

func createChannel(ctx context.Context) (chan OrderNotification, error) {
	orders, err := orders.List(ctx, order.Filter{
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
		items, err := items.List(ctx, o.ID)
		if err != nil {
			return nil, err
		}

		ch <- OrderNotification{
			Order: o,
			Items: items,
		}
	}

	channel[0] = ch
	return ch, nil
}

func sendToChannel(not OrderNotification) {
	ch, ok := channel[0]
	if !ok {
		return
	}

	ch <- not
}
