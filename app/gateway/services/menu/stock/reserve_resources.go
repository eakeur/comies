package stock

import (
	"comies/app/core/entities/ingredient"
	"comies/app/core/workflows/menu"
	"comies/app/core/workflows/stocking"
	"comies/app/sdk/fault"
	"comies/app/sdk/types"
	"context"
)

func (s service) ReserveResources(ctx context.Context, reservationID types.ID, resources ...ingredient.Ingredient) ([]menu.ItemFailed, error) {
	var reservations []stocking.Reservation
	for _, resource := range resources {
		reservations = append(reservations, stocking.Reservation{
			ResourceID: resource.IngredientID,
			Quantity:   resource.Quantity,
		})
	}

	results, err := s.stocks.ReserveResources(ctx, reservationID, reservations)
	if err != nil {
		return nil, fault.Wrap(err)
	}

	var failed []menu.ItemFailed
	for _, it := range results {
		failed = append(failed, menu.ItemFailed{
			ProductID: it.ResourceID,
			Want:      it.Want,
			Got:       it.Got,
			Error:     it.Error,
		})
	}

	return failed, nil
}
