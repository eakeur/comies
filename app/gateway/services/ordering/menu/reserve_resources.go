package menu

import (
	"comies/app/core/entities/ingredient"
	"comies/app/core/workflows/menu"
	"comies/app/core/workflows/ordering"
	"comies/app/sdk/throw"
	"comies/app/sdk/types"
	"context"
)

func (s service) ReserveResources(ctx context.Context, reservationID types.ID, reservation ordering.Reservation) (ordering.Reservation, error) {
	var ignoreInput ingredient.IgnoredList
	for _, ignoring := range reservation.Ignore {
		ignoreInput = append(ignoreInput, types.ID(ignoring))
	}

	replaceInput := ingredient.ReplacedList{}
	for _, r := range reservation.Replace {
		replaceInput[r.From] = r.To
	}

	result, err := s.menu.ReserveProduct(ctx, menu.Reservation{
		ID:        reservation.ID,
		ProductID: reservation.ProductID,
		Quantity:  reservation.Quantity,
		Ignore:    ignoreInput,
		Replace:   replaceInput,
	})
	if err != nil {
		return ordering.Reservation{}, throw.Error(err)
	}

	var failures []ordering.ItemFailed
	for _, f := range result.Failures {
		failures = append(failures, ordering.ItemFailed{
			ProductID: f.ProductID,
			Error:     f.Error,
		})
	}

	reservation.Failures = failures

	return reservation, nil
}
