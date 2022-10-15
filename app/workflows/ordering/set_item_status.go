package ordering

import (
	"comies/app/core/id"
	"comies/app/core/ordering"
	"comies/app/data/items"
	"context"
)

func SetItemStatus(ctx context.Context, id id.ID, status ordering.Status) error {
	if err := ordering.ValidateItemStatus(status); err != nil {
		return err
	}

	return items.SetStatus(ctx, id, status)
}
