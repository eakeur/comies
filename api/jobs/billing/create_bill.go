package billing

import (
	"comies/core/billing/bill"
	"comies/core/billing/item"
	"comies/core/types"
	"context"
	"time"

	"golang.org/x/sync/errgroup"
)

type BillCreation struct {
	Date        time.Time
	Name        types.Text
	ReferenceID types.ID
	Items       []BillItem
}

type BillItem struct {
	ReferenceID types.ID
	Debts       types.Currency
	Credits     types.Currency
	Description types.Text
}

func (j jobs) CreateBill(ctx context.Context, cr BillCreation) (types.ID, error) {

	b, err := bill.Bill{ReferenceID: cr.ReferenceID}.
		WithID(j.createID()).
		WithDate(cr.Date).
		WithName(cr.Name).
		Validate()
	if err != nil {
		return 0, err
	}

	items := make([]item.Item, 0)
	for _, it := range cr.Items {
		id := j.createID()

		save, err := item.Item{
			ID:          id,
			BillID:      b.ID,
			ReferenceID: it.ReferenceID,
			Description: it.Description,
		}.
			WithCredits(it.Credits).
			WithDebts(it.Debts).
			Validate()
		if err != nil {
			return 0, err
		}

		items = append(items, save)
	}

	if len(items) == 0 {
		return 0, bill.ErrMustHaveItems
	}

	err = j.bills.Create(ctx, b)
	if err != nil {
		return 0, err
	}

	eg, ctx := errgroup.WithContext(ctx)
	for _, it := range items {
		it := it

		eg.Go(func() error {
			return j.items.Create(ctx, it)
		})
	}

	if err := eg.Wait(); err != nil {
		return 0, err
	}

	return b.ID, nil
}
