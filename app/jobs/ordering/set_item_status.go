package ordering

import (
	"comies/app/core/ordering"
	"comies/app/core/types"
	"comies/app/data/items"
	"context"
)

func SetItemStatus(ctx context.Context, id types.ID, status ordering.Status) error {
	if err := ordering.ValidateItemStatus(status); err != nil {
		return err
	}

	return items.SetStatus(ctx, id, status)
}
