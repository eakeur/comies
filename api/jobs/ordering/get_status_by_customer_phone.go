package ordering

import (
	"comies/core/ordering/status"
	"context"
)

type Status struct {
	status.Status
	CustomerName string
}

func (w jobs) GetStatusByCustomerPhone(ctx context.Context, phone string) (Status, error) {
	o, err := w.orders.GetByCustomerPhone(ctx, phone)
	if err != nil {
		return Status{}, err
	}

	st, err := w.statuses.GetLatestUpdate(ctx, o.ID)
	if err != nil {
		return Status{}, err
	}

	return Status{
		Status:       st,
		CustomerName: o.CustomerName,
	}, nil
}
