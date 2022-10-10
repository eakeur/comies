package ordering

import (
	"comies/app/core/id"
	"comies/app/core/item"
	"comies/app/core/reservation"
	"comies/app/core/types"
	"comies/app/core/workflows/menu"
	"comies/app/data/items"
	"context"
)

func AddToOrder(ctx context.Context, i item.Item) (ItemAdditionResult, error) {

	if i.OrderID.Empty() {
		return ItemAdditionResult{}, types.ErrMissingID
	}

	id.Create(&i.ID)
	i, err := items.Create(ctx, i)
	if err != nil {
		return ItemAdditionResult{}, err
	}

	res, err := menu.Reserve(ctx, reservation.Reservation{
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
