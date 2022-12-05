package ordering

import (
	"comies/core/ordering/item"
	"comies/core/types"
	"context"
)

func (w jobs) ListItems(ctx context.Context, orderID types.ID) ([]item.Item, error) {
	return w.items.List(ctx, orderID)
}
