package ordering

import (
	"comies/core/ordering/status"
	"context"
)

func (w jobs) CountUnfinishedOrders(ctx context.Context) (status.CountByStatus, error) {
	return w.statuses.CountByStatus(ctx)
}
