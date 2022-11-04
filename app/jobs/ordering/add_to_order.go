package ordering

import (
	"comies/app/core/ordering/item"
	"context"
)

func (w jobs) AddToOrder(ctx context.Context, i item.Item) (item.Item, error) {
	price, err := w.getPrice(ctx, i.ProductID)
	if err != nil {
		return item.Item{}, err
	}

	save, err := i.WithID(w.createID()).WithValue(price).Validate()
	if err != nil {
		return item.Item{}, err
	}

	return save, nil
}
