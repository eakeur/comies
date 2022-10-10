package ordering

import (
	"comies/app/core/id"
	"comies/app/core/item"
	"comies/app/data/items"
	"context"
)

func ListItems(ctx context.Context, orderID id.ID) ([]item.Item, error) {
	list, err := items.List(ctx, orderID)
	if err != nil {
		return nil, err
	}

	return list, err
}
