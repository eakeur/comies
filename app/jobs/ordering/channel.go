package ordering

import (
	"comies/app/core/ordering"
	"comies/app/core/types"
	"comies/app/data/items"
	"comies/app/data/orders"
	"context"
)

type NewOrderNotification struct {
	Order ordering.Order
	Items []ordering.Item
}

type Update struct {
	OrderID types.ID
	Path    string
	Value   interface{}
}

var channel = make(map[types.ID]chan Update)

func Channel(ctx context.Context) (chan Update, error) {
	ch, ok := channel[0]
	if ok {
		return ch, nil
	}

	list, err := orders.List(ctx, ordering.OrderFilter{Status: ordering.UndoneOrderStatuses})
	if err != nil {
		return nil, err
	}

	ch = make(chan Update)
	for _, o := range list {
		i, err := items.List(ctx, o.ID)
		if err != nil {
			return nil, err
		}

		ch <- Update{
			OrderID: o.ID,
			Path:    "*",
			Value: NewOrderNotification{
				Order: o,
				Items: i,
			},
		}
	}

	channel[0] = ch
	return ch, nil
}

func sendch(orderID types.ID, path string, val interface{}) {
	go func(orderID types.ID, path string, val interface{}) {
		if ch, ok := channel[0]; ok {
			ch <- Update{
				OrderID: orderID,
				Path:    path,
				Value:   val,
			}
		}
	}(orderID, path, val)
}

const (
	newOrderPath     = "*"
	changeStatusPath = "order.status"
)
