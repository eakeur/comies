package ordering

import (
	"comies/app/core/ordering"
	"comies/app/core/types"
	"comies/app/data/ids"
	"comies/app/data/items"
	"context"
)

func AddToOrder(ctx context.Context, i ordering.Item) (types.ID, error) {
	i.ID = ids.Create()
	if err := ordering.ValidateItem(i); err != nil {
		return 0, err
	}

	return i.ID, items.Create(ctx, i)
}
