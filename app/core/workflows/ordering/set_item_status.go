package ordering

import (
	"comies/app/core/id"
	"comies/app/core/item"
	"comies/app/data/items"
	"context"
)

func SetItemStatus(ctx context.Context, id id.ID, status item.Status) error {
	err := items.SetStatus(ctx, id, status)
	if err != nil {
		return err
	}

	return nil
}
