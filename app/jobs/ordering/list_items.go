package ordering

import (
	"comies/app/core/ordering/item"
	"comies/app/core/types"
	"context"
)

func (w jobs) ListItems(ctx context.Context, orderID types.ID) ([]item.Item, error) {
	return w.items.List(ctx, orderID)
}
