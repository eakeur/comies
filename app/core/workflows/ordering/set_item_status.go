package ordering

import (
	"comies/app/core/entities/item"
	"comies/app/core/throw"
	"comies/app/core/types"
	"context"
)

func (w workflow) SetItemStatus(ctx context.Context, id types.ID, status item.Status) error {
	err := w.items.SetStatus(ctx, id, status)
	if err != nil {
		return throw.Error(err)
	}

	return nil
}
