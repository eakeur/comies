package ordering

import (
	"comies/app/core/ordering"
	"comies/app/core/types"
	"comies/app/data/items"
	"comies/app/data/orders"
	"context"
	"time"
)

type OrderConfirmation struct {
	OrderID         types.ID
	DeliveryType    ordering.Type
	CustomerName    string
	CustomerAddress string
	CustomerPhone   string
}

func PlaceOrder(ctx context.Context, c OrderConfirmation) (ordering.Order, []ordering.Item, error) {
	o, err := orders.GetByID(ctx, c.OrderID)
	if err != nil {
		return ordering.Order{}, nil, err
	}

	o.DeliveryType = c.DeliveryType
	o.PlacedAt = time.Now().UTC()
	o.Customer.Name = c.CustomerName
	o.Customer.Phone = c.CustomerPhone
	o.Customer.Address = c.CustomerAddress

	if err := ordering.CheckIfOrderIsPlaceable(o); err != nil {
		return ordering.Order{}, nil, err
	}

	itemsList, err := items.List(ctx, o.ID)
	if err != nil || len(itemsList) <= 0 {
		return ordering.Order{}, nil, ordering.ErrInvalidNumberOfItems
	}

	if err := updateOrderOnConfirmation(ctx, o); err != nil {
		return ordering.Order{}, nil, err
	}

	if err = orders.UpdateFlow(ctx, ordering.NewOrderFlow(o)); err != nil {
		return ordering.Order{}, nil, err
	}

	return o, itemsList, nil

}

func updateOrderOnConfirmation(ctx context.Context, o ordering.Order) error {
	if err := orders.SetDeliveryType(ctx, o.ID, o.DeliveryType); err != nil {
		return err
	}

	if err := orders.SetPlacedAt(ctx, o.ID, o.PlacedAt); err != nil {
		return err
	}

	if err := orders.SetObservation(ctx, o.ID, o.Observations); err != nil {
		return err
	}

	return nil
}
