package menu

import (
	"comies/app/core/menu"
	"comies/app/data/ingredients"
	"context"
	"errors"
	"sync"
)

func Reserve(ctx context.Context, r menu.Reservation) (failures []menu.ReservationFailure, err error) {
	ings, err := ingredients.List(ctx, r.ProductID)
	if err != nil {
		return nil, err
	}

	if len(ings) == 0 {
		_, err = CreateMovement(ctx, menu.ReservationToReservedMovement(r), r.Quantity)
		if err != nil && errors.Is(err, menu.ErrStockNegative) {
			failures = append(failures, menu.ReservationFailure{
				For:       r.ReserveFor,
				ProductID: r.ProductID,
				Error:     err,
			})
		}

		return
	}

	return reserveIngredients(ctx, r, ings)
}

func reserveIngredients(ctx context.Context, r menu.Reservation, ings []menu.Ingredient) (failures []menu.ReservationFailure, err error) {
	var wg = sync.WaitGroup{}
	var errs = make(chan error, len(ings))

	for _, ing := range menu.ModifyIngredientsList(ings, r.Specifics, r.Quantity) {
		wg.Add(1)

		go func(ctx context.Context, i menu.Ingredient) {
			defer wg.Done()
			f, err := Reserve(ctx, menu.ReservationToIngredientReservation(r, i))
			if err != nil {
				errs <- err
			}

			failures = append(failures, f...)
		}(ctx, ing)
	}

	wg.Wait()

	if len(errs) > 0 {
		return nil, <-errs
	}

	return
}
