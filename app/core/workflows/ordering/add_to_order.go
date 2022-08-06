package ordering

import (
	"comies/app/core/entities/item"
	"comies/app/core/entities/reservation"
	"comies/app/core/throw"
	"context"
)

func (w workflow) AddToOrder(ctx context.Context, i item.Item) (ItemAdditionResult, error) {

	if i.OrderID.Empty() {
		return ItemAdditionResult{}, throw.Error(throw.ErrMissingID)
	}

	w.id.Create(&i.ID)
	i, err := w.items.Create(ctx, i)
	if err != nil {
		return ItemAdditionResult{}, throw.Error(err)
	}

	res, err := w.products.Reserve(ctx, reservation.Reservation{
		ID:        i.ID,
		ProductID: i.ProductID,
		Quantity:  i.Quantity,
		Ignore:    i.Details.IgnoreIngredients,
		Replace:   i.Details.ReplaceIngredients,
	})
	if err != nil {
		return ItemAdditionResult{}, throw.Error(err)
	}

	result := ItemAdditionResult{}
	if len(res) > 0 {
		result.Failed = res
	}

	return result, nil
}
