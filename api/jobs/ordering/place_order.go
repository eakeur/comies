package ordering

import (
	"comies/core/ordering/item"
	"comies/core/ordering/order"
	"comies/core/ordering/status"
	"comies/core/types"
	"comies/jobs/billing"
	"comies/jobs/menu"
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

	billItems := make([]billing.BillItem, len(conf.items))
	eg, ctx := errgroup.WithContext(ctx)
	for i, item := range conf.items {
		item := item
		i := i
		eg.Go(func() error {
			save, err := w.createPricedItem(ctx, item)
			if err != nil {
				return err
			}

			billItems[i] = billing.BillItem{
				ReferenceID: save.ID,
				Credits:     save.Value,
			}

			return w.menu.DispatchProduct(ctx, menu.Dispatcher{
				ProductID: save.ProductID,
				Price:     save.Value,
				AgentID:   save.ID,
				Quantity:  save.Quantity,
				Date:      conf.Time,
			})
		})
	}

	if err := eg.Wait(); err != nil {
		return order.Order{}, err
	}

	eg.Go(func() error {
		st, err := status.Status{OrderID: o.ID}.
			WithOccurredAt(conf.Time).
			WithValue(status.PreparingStatus).
			Validate()
		if err != nil {
			return err
		}

		return w.statuses.Update(ctx, st)
	})

	return o, w.createOrderBill(ctx, o, billItems)
}

func (w jobs) createPricedItem(ctx context.Context, i item.Item) (item.Item, error) {
	price, err := w.menu.GetProductLatestPriceByID(ctx, i.ProductID)
	if err != nil {
		return item.Item{}, err
	}

	save, err := i.WithValue(price).Validate()
	if err != nil {
		return item.Item{}, err
	}

	err = w.items.Create(ctx, save)
	if err != nil {
		return item.Item{}, err
	}

	return save, nil
}

func (w jobs) createOrderBill(ctx context.Context, o order.Order, items []billing.BillItem) error {
	_, err := w.billing.CreateBill(ctx, billing.BillCreation{
		ReferenceID: o.ID,
		Date:        o.PlacedAt,
		Items:       items,
	})
	if err != nil {
		return err
	}
	return nil
}
