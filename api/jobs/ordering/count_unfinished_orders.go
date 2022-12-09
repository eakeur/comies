package ordering

import (
	"comies/core/types"
	"context"
)

func (w jobs) CountOrdersByStatus(ctx context.Context, status types.Status) (types.Quantity, error) {
	return w.statuses.CountByStatus(ctx, status)
}
