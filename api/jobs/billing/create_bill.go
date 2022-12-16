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

type BillSummary struct {
	ID  types.ID
	Sum types.Amount
}

func (j jobs) CreateBill(ctx context.Context, cr BillCreation) (BillSummary, error) {

	if len(cr.Items) == 0 {
		return BillSummary{}, bill.ErrMustHaveItems
	}

	b, err := bill.Bill{ReferenceID: cr.ReferenceID}.
		WithID(j.createID()).
		WithDate(cr.Date).
		WithName(cr.Name).
		Validate()
	if err != nil {
		return BillSummary{}, err
	}

	err = j.bills.Create(ctx, b)
	if err != nil {
		return BillSummary{}, err
	}

	eg, cctx := errgroup.WithContext(ctx)
	for _, it := range cr.Items {
		it := it

		eg.Go(func() error {
			save, err := item.Item{
				ID:          j.createID(),
				BillID:      b.ID,
				ReferenceID: it.ReferenceID,
				Description: it.Description,
			}.
				WithCredits(it.Credits).
				WithDebts(it.Debts).
				Validate()
			if err != nil {
				return err
			}

			return j.items.Create(cctx, save)
		})

	}

	if err := eg.Wait(); err != nil {
		return BillSummary{}, err
	}

	sum, err := j.items.SumByBillID(ctx, b.ID)
	if err != nil {
		return BillSummary{}, err
	}

	return BillSummary{
		ID: b.ID,
		Sum: types.Amount{
			Value:    sum,
			Net:      sum,
			Currency: "BRL",
		},
	}, nil
}
