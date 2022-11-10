package ordering

import (
	"comies/app/core/ordering/item"
	"comies/app/core/ordering/order"
	"comies/app/core/ordering/status"
	"comies/app/core/types"
	"context"
	"time"
)

type Order struct {
	items           []item.Item
	DeliveryType    types.Type
	Observations    string
	CustomerName    string
	CustomerPhone   string
	CustomerAddress string
	Time            time.Time
}

func (w jobs) PlaceOrder(ctx context.Context, conf Order) (order.Order, error) {
	if len(conf.items) <= 0 {
		return order.Order{}, order.ErrInvalidNumberOfItems
	}

	o, err := order.Order{}.
		WithID(w.createID()).
		WithPlacedAt(conf.Time).
		WithDeliveryType(conf.DeliveryType).
		WithCustomer(conf.CustomerName, conf.CustomerPhone, conf.CustomerAddress).
		WithObservations(conf.Observations).
		Validate()
	if err != nil {
		return order.Order{}, err
	}

	err = w.orders.Create(ctx, o)
	if err != nil {
		return order.Order{}, err
	}

	for _, item := range conf.items {
		_, err = item.Validate()
		if err != nil {
			return order.Order{}, err
		}

		err = w.items.Create(ctx, item)
		if err != nil {
			return order.Order{}, err
		}

		err := w.dispatchProduct(ctx, item.ProductID, item.ID, item.Quantity)
		if err != nil {
			return order.Order{}, err
		}
	}

	st, err := status.Status{OrderID: o.ID}.
		WithOccurredAt(conf.Time).
		WithValue(status.PreparingStatus).
		Validate()
	if err != nil {
		return order.Order{}, err
	}

	err = w.statuses.Update(ctx, st)
	if err != nil {
		return order.Order{}, err
	}

	return o, nil
}
