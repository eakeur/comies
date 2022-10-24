package ordering

import (
	"comies/app/core/entities/item"
	"comies/app/core/types"
	"context"
)

func (w jobs) ListItems(ctx context.Context, orderID types.ID) ([]item.Item, error) {
	list, err := w.items.List(ctx, orderID)
	if err != nil {
		return nil, err
	}

	return list, err
}
