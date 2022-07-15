package ordering

import (
	"comies/app/core/entities/item"
	"comies/app/sdk/throw"
	"comies/app/sdk/types"
	"context"
)

func (w workflow) SetItemStatus(ctx context.Context, id types.ID, status item.Status) error {
	err := w.items.SetStatus(ctx, id, status)
	if err != nil {
		return throw.Error(err)
	}

	return nil
}
