package ordering

import (
	"comies/app/core/entities/item"
	"comies/app/sdk/throw"
	"comies/app/sdk/types"
	"context"
)

func (w workflow) ListItems(ctx context.Context, orderID types.ID) ([]item.Item, error) {
	list, err := w.items.List(ctx, orderID)
	if err != nil {
		return nil, throw.Error(err)
	}

	return list, err
}
