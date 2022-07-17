package menu

import (
	"comies/app/core/entities/ingredient"
	"comies/app/core/entities/movement"
	"comies/app/sdk/throw"
	"comies/app/sdk/types"
	"context"
	"sync"
	"time"
)

func (w workflow) ReserveProduct(ctx context.Context, r Reservation) (Reservation, error) {

	var (
		params = map[string]interface{}{
			"reservation_id": r.ID.String(),
			"product_id":     r.ProductID.String(),
		}
	)

	ingredients, err := w.ingredients.List(ctx, r.ProductID)
	if err != nil {
		return Reservation{}, throw.Error(err).Params(params)
	}

	if len(ingredients) == 0 {
		failures, err := w.createReservedMovements(ctx, r.ID, ingredient.Ingredient{
			Quantity:     r.Quantity,
			IngredientID: r.ProductID,
		})
		if err != nil {
			return Reservation{}, throw.Error(err).Params(params)
		}
		r.Failures = failures

		return r, nil
	}

	failures, err := w.createReservedMovements(ctx, r.ID, ingredient.IgnoreAndReplace(
		ingredients, r.Ignore, r.Replace,
		func(i ingredient.Ingredient) ingredient.Ingredient {
			i.Quantity *= r.Quantity
			return i
		},
	)...)

	r.Failures = failures

	return r, nil
}

func (w workflow) createReservedMovements(ctx context.Context, reservationID types.ID, ingredients ...ingredient.Ingredient) ([]ItemFailed, error) {

	var (
		errors   = make(chan error, len(ingredients))
		failures []ItemFailed
		wg       sync.WaitGroup
	)

	for _, ing := range ingredients {
		ing := ing
		wg.Add(1)

		go func() {
			defer wg.Done()
			if _, err := w.CreateMovement(ctx, movement.Movement{
				ProductID: ing.IngredientID,
				Type:      movement.ReservedMovement,
				Date:      time.Now(),
				Quantity:  ing.Quantity,
				AgentID:   reservationID,
			}); err != nil {
				errors <- throw.Error(err)
				failures = append(failures, ItemFailed{
					ProductID: ing.IngredientID,
					Error:     err,
				})
			}
		}()
	}
	wg.Wait()

	if len(errors) > 0 {
		return nil, throw.Error(<-errors)
	}

	return failures, nil
}
