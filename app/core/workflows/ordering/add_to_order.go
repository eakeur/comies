package ordering

import (
	"context"
	"gomies/app/core/entities/item"
	"gomies/app/sdk/fault"
)

func (w workflow) AddToOrder(ctx context.Context, i item.Item) (ItemAdditionResult, error) {

	if i.OrderID.Empty() {
		return ItemAdditionResult{}, fault.Wrap(fault.ErrMissingID)
	}

	i, err := w.items.Create(ctx, i)
	if err != nil {
		return ItemAdditionResult{}, fault.Wrap(err)
	}

	res, err := w.products.ReserveResources(ctx, i.ID, Reservation{
		ID:        i.ID,
		ProductID: i.ProductID,
		Quantity:  i.Quantity,
		Ignore:    i.Details.IgnoreIngredients,
		Replace:   i.Details.ReplaceIngredients,
	})
	if err != nil {
		return ItemAdditionResult{}, fault.Wrap(err)
	}

	result := ItemAdditionResult{Item: i}
	if len(res.Failures) > 0 {
		result.Failed = []Reservation{res}
	} else {
		result.Succeeded = []Reservation{res}
	}

	return result, nil
}
