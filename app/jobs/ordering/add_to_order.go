package ordering

import (
	reservation "comies/app/core/menu"
	"comies/app/core/ordering"
	"comies/app/core/types"
	"comies/app/data/ids"
	"comies/app/data/items"
	"comies/app/jobs/menu"
	"context"
)

func AddToOrder(ctx context.Context, i ordering.Item) (failure []reservation.ReservationFailure, id types.ID, err error) {
	i.ID = ids.Create()
	if err := ordering.ValidateItem(i); err != nil {
		return nil, 0, err
	}

	if res, err := menu.Reserve(ctx, reservation.Reservation{
		ID:        i.ID,
		ProductID: i.ProductID,
		Quantity:  i.Quantity,
		Specifics: i.Specification,
	}); err != nil || len(res) > 0 {
		return res, 0, err
	}

	return nil, i.ID, items.Create(ctx, i)
}
