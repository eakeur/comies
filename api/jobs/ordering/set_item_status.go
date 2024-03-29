package ordering

import (
	"comies/core/types"
	"context"
)

func (w jobs) SetItemStatus(ctx context.Context, id types.ID, status types.Status) error {
	return w.items.SetStatus(ctx, id, status)
}
