package ordering

import (
	"comies/core/ordering/item"
	"comies/core/ordering/order"
	"comies/core/ordering/status"
	"comies/core/types"
	"context"
	"time"

	"golang.org/x/sync/errgroup"
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

	eg, ctx := errgroup.WithContext(ctx)
	for _, item := range conf.items {
		item := item
		eg.Go(func() error {
			price, err := w.getPrice(ctx, item.ProductID)
			if err != nil {
				return err
			}

			save, err := item.WithValue(price).Validate()
			if err != nil {
				return err
			}

			err = w.items.Create(ctx, save)
			if err != nil {
				return err
			}

			return w.dispatchProduct(ctx, types.Dispatcher{
				ProductID: save.ProductID,
				Price:     save.Value,
				AgentID:   save.ID,
				Quantity:  save.Quantity,
			})
		})
	}

	if err := eg.Wait(); err != nil {
		return order.Order{}, err
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
