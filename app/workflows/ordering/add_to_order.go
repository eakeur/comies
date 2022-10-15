package ordering

import (
	reservation "comies/app/core/menu"
	"comies/app/core/ordering"
	"comies/app/data/ids"
	"comies/app/data/items"
	"comies/app/workflows/menu"
	"context"
)

func AddToOrder(ctx context.Context, i ordering.Item) (failure []reservation.ReservationFailure, err error) {
	i.ID = ids.Create()
	if err := ordering.ValidateItem(i); err != nil {
		return nil, err
	}

	res, err := menu.Reserve(ctx, reservation.Reservation{
		ID:        i.ID,
		ProductID: i.ProductID,
		Quantity:  i.Quantity,
		Specifics: reservation.IngredientSpecification(i.Specification),
	})
	if err != nil || len(res) > 0 {
		return res, err
	}

	return nil, items.Create(ctx, i)
}
