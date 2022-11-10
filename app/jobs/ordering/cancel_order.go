package ordering

import (
	"comies/app/core/ordering/order"
	"comies/app/core/ordering/status"
	"comies/app/core/types"
	"context"
	"time"
)

func (w jobs) CancelOrder(ctx context.Context, id types.ID) error {

	o, err := w.statuses.GetLastUpdate(ctx, id)
	if err != nil {
		return err
	}

	if o.Value > status.PreparingStatus {
		return order.ErrAlreadyPrepared
	}

	return w.statuses.Update(ctx, status.Status{OrderID: id}.WithOccurredAt(time.Now()).WithValue(status.CanceledStatus))
}
