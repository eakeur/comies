package ordering

import (
	"comies/app/core/types"
	"context"
)

func (w jobs) SetItemStatus(ctx context.Context, id types.ID, status types.Status) error {
	err := w.items.SetStatus(ctx, id, status)
	if err != nil {
		return err
	}

	return nil
}
