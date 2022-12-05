package ordering

import (
	"comies/core/ordering/order"
	"comies/core/ordering/status"
	"comies/core/types"
	"context"
	"time"
)

func (w jobs) CancelOrder(ctx context.Context, id types.ID) error {

	o, err := w.statuses.GetLatestUpdate(ctx, id)
	if err != nil {
		return err
	}

	if o.Value > status.PreparingStatus {
		return order.ErrAlreadyPrepared
	}

	return w.statuses.Update(ctx, status.Status{OrderID: id}.WithOccurredAt(time.Now()).WithValue(status.CanceledStatus))
}
