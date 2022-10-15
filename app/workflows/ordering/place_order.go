package ordering

import (
	"comies/app/core/id"
	"comies/app/core/ordering"
	"comies/app/data/items"
	"comies/app/data/orders"
	"comies/app/workflows/menu"
	"context"
	"time"
)

type OrderConfirmation struct {
	OrderID      id.ID
	DeliveryType ordering.Type
	CustomerName string
	CustomerAddress string
	CustomerPhone string
}

func PlaceOrder(ctx context.Context, c OrderConfirmation) (ordering.Order, error) {
	o, err := orders.GetByID(ctx, c.OrderID)
	if err != nil {
		return ordering.Order{}, err
	}

	o.DeliveryType = c.DeliveryType
	o.PlacedAt = time.Now().UTC()
	o.Customer.Name = c.CustomerName
	o.Customer.Phone = c.CustomerPhone
	o.Customer.Address = c.CustomerAddress

	if err := ordering.CheckIfOrderIsPlaceable(o); err != nil {
		return ordering.Order{}, err
	}

	itemsList, err := items.List(ctx, o.ID)
	if err != nil || len(itemsList) <= 0{
		return ordering.Order{}, ordering.ErrInvalidNumberOfItems
	}

	if err := orders.SetDeliveryMode(ctx, o.ID, o.DeliveryType); err != nil {
		return ordering.Order{}, err
	}

	if err = orders.UpdateFlow(ctx, ordering.NewOrderFlow(o)); err != nil {
		return ordering.Order{}, err
	}

	sendToChannel(o.ID, "*", NewOrderNotification{
		Order: o,
		Items: itemsList,
	})

	go func() {
		for _, item := range itemsList {
			menu.UpdateReservation(ctx, item.ID, true)
		}
	}()

	return o, nil

}
