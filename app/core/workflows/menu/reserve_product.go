package menu

import (
	"comies/app/core/ingredient"
	"comies/app/core/movement"
	"comies/app/core/product"
	"comies/app/core/reservation"
	"comies/app/data/ingredients"
	"context"
	"errors"
	"sync"
)

func Reserve(ctx context.Context, r reservation.Reservation) (failures []reservation.Failure, err error) {
	ingredients, err := ingredients.List(ctx, r.ProductID)
	if err != nil {
		return nil, err
	}

	wg := sync.WaitGroup{}
	errs := make(chan error, len(ingredients))

	if len(ingredients) == 0 {
		if _, err := CreateMovement(ctx, movement.Movement{
			ProductID: r.ProductID,
			Quantity:  r.Quantity,
			AgentID:   r.ID,
			Type:      movement.ReservedType,
		}); err != nil {
			if errors.Is(err, product.ErrStockNegative) {
				failures = append(failures, reservation.Failure{
					For:       r.ReserveFor,
					ProductID: r.ProductID,
					Error:     err,
				})
			}

			return nil, err
		}
	} else {
		for _, ing := range ingredient.IgnoreAndReplace(ingredients, r.Ignore, r.Replace, func(i ingredient.Ingredient) ingredient.Ingredient {
			i.Quantity *= r.Quantity
			return i
		}) {
			ing := ing
			wg.Add(1)
			go func() {
				defer wg.Done()
				f, err := Reserve(ctx, reservation.Reservation{
					ID:         r.ID,
					ReserveFor: ing.ProductID,
					ProductID:  ing.IngredientID,
					Quantity:   ing.Quantity,
					Ignore:     r.Ignore,
					Replace:    r.Replace,
				})
				if err != nil {
					errs <- err
				}

				failures = append(failures, f...)
			}()

		}
		wg.Wait()
	}

	if len(errs) > 0 {
		return nil, <-errs
	}

	return failures, nil
}
