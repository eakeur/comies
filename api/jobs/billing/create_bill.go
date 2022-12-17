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
	Name        string
	ReferenceID types.ID
	Items       []BillItem
}

type BillItem struct {
	Name        string
	ReferenceID types.ID
	UnitPrice   types.Currency
	Quantity    types.Quantity
	Discounts   types.Currency
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
		WithName(types.Text(cr.Name)).
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
				Name:        it.Name,
				UnitPrice:   it.UnitPrice,
				Quantity:    it.Quantity,
				Discounts:   it.Discounts,
			}.Validate()
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
		ID:  b.ID,
		Sum: sum,
	}, nil
}
