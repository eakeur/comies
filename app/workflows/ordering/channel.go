package ordering

import (
	"comies/app/core/id"
	"comies/app/core/ordering"
	"comies/app/data/items"
	"comies/app/data/orders"
	"context"
)

type NewOrderNotification struct {
	Order ordering.Order
	Items []ordering.Item
}

type Update struct {
	OrderID id.ID
	Path string
	Value interface{}
}

var channel = make(map[id.ID]chan Update)

func Channel(ctx context.Context) (chan Update, error) {
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

func createChannel(ctx context.Context) (chan Update, error) {
	list, err := orders.List(ctx, ordering.UndoneOrderStatuses)
	if err != nil {
		return nil, err
	}

	ch := make(chan Update)
	for _, o := range list {
		i, err := items.List(ctx, o.ID)
		if err != nil {
			return nil, err
		}

		ch <- Update{
			OrderID: o.ID,
			Path: "*",
			Value: NewOrderNotification{
				Order: o,
				Items: i,
			},
		}
	}

	channel[0] = ch
	return ch, nil
}

func sendToChannel(orderID id.ID, path string, val interface{}) {
	go func(){
		ch, ok := channel[0]
		if !ok {
			return
		}

		ch <- Update{
			OrderID: orderID,
			Path:    path,
			Value:   val,
			}
	}()
}
