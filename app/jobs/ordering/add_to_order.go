package ordering

import (
	"comies/app/core/ordering/item"
	"context"
)

func (w jobs) AddToOrder(ctx context.Context, i item.Item) {
	save, err := i.WithID().Validate()
	if err != nil {
		return err
	}

	w.id.Create(&i.ID)
	i, err := w.items.Create(ctx, i)
	if err != nil {
		return ItemAdditionResult{}, err
	}

	res, err := w.products.Reserve(ctx, reservation.Reservation{
		ID:        i.ID,
		ProductID: i.ProductID,
		Quantity:  i.Quantity,
		Ignore:    i.Details.IgnoreIngredients,
		Replace:   i.Details.ReplaceIngredients,
	})
	if err != nil {
		return ItemAdditionResult{}, err
	}

	result := ItemAdditionResult{}
	if len(res) > 0 {
		result.Failed = res
	}

	return result, nil
}
