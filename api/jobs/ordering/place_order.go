package ordering

import (
	"comies/core/ordering/item"
	"comies/core/ordering/order"
	"comies/core/ordering/status"
	"comies/core/types"
	"comies/jobs/billing"
	"comies/jobs/menu"
	"context"
	"fmt"
	"time"

	"golang.org/x/sync/errgroup"
)

type Order struct {
	Items           []item.Item
	DeliveryType    types.Type
	Observations    string
	CustomerName    string
	CustomerPhone   string
	CustomerAddress string
	Date            time.Time
}

func (w jobs) PlaceOrder(ctx context.Context, conf Order) (order.Order, error) {
	if len(conf.Items) <= 0 {
		return order.Order{}, order.ErrInvalidNumberOfItems
	}

	o, err := order.Order{}.
		WithID(w.createID()).
		WithPlacedAt(conf.Date).
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

	billItems := make([]billing.BillItem, len(conf.Items))
	eg, cctx := errgroup.WithContext(ctx)
	for i, item := range conf.Items {
		item := item
		i := i
		eg.Go(func() error {
			save, err := w.createPricedItem(cctx, item.WithOrderID(o.ID))
			if err != nil {
				return err
			}

			name, err := w.menu.GetProductNameByID(ctx, save.ProductID)
			if err != nil {
				return err
			}

			billItems[i] = billing.BillItem{
				ReferenceID: save.ID,
				Credits:     save.Price,
				Description: types.Text(name),
			}

			return w.menu.DispatchProduct(cctx, menu.Dispatcher{
				ProductID: save.ProductID,
				Price:     save.Price,
				AgentID:   save.ID,
				Quantity:  save.Quantity,
				Date:      conf.Date,
			})
		})
	}

	if err := eg.Wait(); err != nil {
		return order.Order{}, err
	}

	st, err := status.Status{OrderID: o.ID}.
		WithOccurredAt(conf.Date).
		WithValue(status.PreparingStatus).
		Validate()
	if err != nil {
		return o, err
	}

	if err := w.statuses.Update(ctx, st); err != nil {
		return o, err
	}

	if err := w.createOrderBill(ctx, o, billItems); err != nil {
		return o, err
	}

	return o, nil
}

func (w jobs) createPricedItem(ctx context.Context, i item.Item) (item.Item, error) {
	price, err := w.menu.GetProductLatestPriceByID(ctx, i.ProductID)
	if err != nil {
		return item.Item{}, err
	}

	save, err := i.WithID(w.createID()).WithValue(price).
		WithStatus(item.PreparingItemStatus).
		Validate()
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
		Name:        types.Text(fmt.Sprintf("Conta de %s", o.CustomerName)),
		ReferenceID: o.ID,
		Date:        o.PlacedAt,
		Items:       items,
	})
	if err != nil {
		return err
	}
	return nil
}
