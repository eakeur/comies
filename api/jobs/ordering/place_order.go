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

type Ticket struct {
	DeliveryType    types.Type
	Observations    string
	CustomerName    string
	CustomerPhone   string
	CustomerAddress string
	Date            time.Time
	Items           []TicketItem
}

type TicketItem struct {
	ProductID    types.ID
	Quantity     types.Quantity
	Observations string
}

type OrderSummary struct {
	Items         []item.Item  `json:"items"`
	BillID        types.ID     `json:"bill_id"`
	BillAmountDue types.Amount `json:"bill_amount_due"`
	order.Order
}

func (w jobs) PlaceOrder(ctx context.Context, tk Ticket) (OrderSummary, error) {
	if len(tk.Items) <= 0 {
		return OrderSummary{}, order.ErrInvalidNumberOfItems
	}

	ord, err := w.createOrderFromTicket(ctx, tk)
	if err != nil {
		return OrderSummary{}, err
	}

	items, err := w.createOrderItems(ctx, ord, tk.Items)
	if err != nil {
		return OrderSummary{}, err
	}

	bill, err := w.createOrderBill(ctx, ord, items)
	if err != nil {
		return OrderSummary{}, err
	}

	st, err := status.Status{OrderID: ord.ID}.
		WithOccurredAt(tk.Date).
		WithValue(status.PreparingStatus).
		Validate()
	if err != nil {
		return OrderSummary{}, err
	}

	if err := w.statuses.Update(ctx, st); err != nil {
		return OrderSummary{}, err
	}

	return OrderSummary{
		Items:         items,
		BillID:        bill.ID,
		BillAmountDue: bill.Sum,
		Order:         ord,
	}, nil
}

func (w jobs) createOrderFromTicket(ctx context.Context, tk Ticket) (order.Order, error) {
	o, err := order.Order{}.
		WithID(w.createID()).
		WithPlacedAt(tk.Date).
		WithDeliveryType(tk.DeliveryType).
		WithCustomer(tk.CustomerName, tk.CustomerPhone, tk.CustomerAddress).
		WithObservations(tk.Observations).
		Validate()
	if err != nil {
		return order.Order{}, err
	}

	return o, w.orders.Create(ctx, o)
}

func (w jobs) createOrderItems(ctx context.Context, o order.Order, tki []TicketItem) ([]item.Item, error) {
	eg, ctx := errgroup.WithContext(ctx)
	items := make([]item.Item, len(tki))

	for i, it := range tki {
		it := it
		i := i

		eg.Go(func() error {
			price, err := w.menu.GetProductLatestPriceByID(ctx, it.ProductID)
			if err != nil {
				return err
			}

			save, err := item.Item{
				ProductID:    it.ProductID,
				Quantity:     it.Quantity,
				Observations: it.Observations,
			}.
				WithID(w.createID()).
				WithOrderID(o.ID).
				WithValue(price).
				WithStatus(item.PreparingItemStatus).
				Validate()
			if err != nil {
				return err
			}

			err = w.items.Create(ctx, save)
			if err != nil {
				return err
			}

			items[i] = save

			return w.menu.DispatchProduct(ctx, menu.Dispatcher{
				ProductID: save.ProductID,
				Price:     save.Price,
				AgentID:   save.ID,
				Quantity:  save.Quantity,
				Date:      o.PlacedAt,
			})
		})
	}

	if err := eg.Wait(); err != nil {
		return nil, err
	}

	return items, nil
}

func (w jobs) createOrderBill(ctx context.Context, o order.Order, orderItems []item.Item) (billing.BillSummary, error) {

	eg, cctx := errgroup.WithContext(ctx)
	items := make([]billing.BillItem, len(orderItems))

	for i, it := range orderItems {
		it := it
		i := i

		eg.Go(func() error {
			name, err := w.menu.GetProductNameByID(cctx, it.ProductID)
			if err == nil {
				items[i] = billing.BillItem{
					ReferenceID: it.ProductID,
					Credits:     it.Price * types.Currency(it.Quantity),
					Description: types.Text(fmt.Sprintf("%d - %s", it.Quantity, name)),
				}
			}

			return err
		})
	}

	if err := eg.Wait(); err != nil {
		return billing.BillSummary{}, err
	}

	summ, err := w.billing.CreateBill(ctx, billing.BillCreation{
		Name:        types.Text(fmt.Sprintf("Conta de %s", o.CustomerName)),
		ReferenceID: o.ID,
		Date:        o.PlacedAt,
		Items:       items,
	})
	if err != nil {
		return billing.BillSummary{}, err
	}

	return summ, nil
}
