package ordering

import (
	"comies/app/core/ordering/status"
	"context"
)

func (w jobs) CountUnfinishedOrders(ctx context.Context) (status.CountByStatus, error) {
	return w.statuses.CountByStatus(ctx)
}
