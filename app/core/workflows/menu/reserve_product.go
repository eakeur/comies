package menu

import (
	"comies/app/core/entities/ingredient"
	"comies/app/core/entities/movement"
	"comies/app/core/entities/product"
	"comies/app/core/entities/reservation"
	"comies/app/sdk/throw"
	"context"
	"errors"
	"sync"
)

func (w workflow) Reserve(ctx context.Context, r reservation.Reservation) (failures []reservation.Failure, err error) {

	var (
		params = map[string]interface{}{
			"reservation_id": r.ID.String(),
			"product_id":     r.ProductID.String(),
		}
	)

	ingredients, err := w.ingredients.List(ctx, r.ProductID)
	if err != nil {
		return nil, throw.Error(err).Params(params)
	}

	wg := sync.WaitGroup{}
	errs := make(chan error, len(ingredients))

	if len(ingredients) == 0 {
		if _, err := w.CreateMovement(ctx, movement.Movement{
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

			return nil, throw.Error(err)
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
				f, err := w.Reserve(ctx, reservation.Reservation{
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
		return nil, throw.Error(<-errs).Params(params)
	}

	return failures, nil
}
