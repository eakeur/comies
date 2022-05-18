package ordering

import (
	"context"
	"gomies/app/core/entities/item"
	"gomies/app/sdk/fault"
	"gomies/app/sdk/types"
)

func (w workflow) SetItemStatus(ctx context.Context, id types.ID, status item.Status) error {
	err := w.items.SetStatus(ctx, id, status)
	if err != nil {
		return fault.Wrap(err)
	}

	return nil
}
