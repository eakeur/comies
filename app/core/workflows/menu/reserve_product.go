package menu

import (
	"comies/app/core/entities/ingredient"
	"comies/app/core/entities/movement"
	"comies/app/core/entities/product"
	"comies/app/core/entities/reservation"
	"comies/app/sdk/throw"
	"comies/app/sdk/types"
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
	creator := func(productID types.ID, quantity types.Quantity) {
		defer wg.Done()
		if _, err := w.CreateMovement(ctx, movement.Movement{
			ProductID: productID,
			Quantity:  quantity,
			AgentID:   r.ID,
			Type:      movement.ReservedType,
		}); err != nil {
			if errors.Is(err, product.ErrStockNegative) {
				failures = append(failures, reservation.Failure{ProductID: r.ProductID, Error: err})
				err = nil
			}

			errs <- throw.Error(err)
		}
	}

	if len(ingredients) == 0 {
		wg.Add(1)
		creator(r.ProductID, r.Quantity)
	} else {
		for _, ing := range ingredient.IgnoreAndReplace(ingredients, r.Ignore, r.Replace, func(i ingredient.Ingredient) ingredient.Ingredient {
			i.Quantity *= r.Quantity
			return i
		}) {
			ing := ing
			wg.Add(1)
			go creator(ing.IngredientID, ing.Quantity)
		}
		wg.Wait()
	}

	if len(errs) > 0 {
		return nil, throw.Error(<-errs).Params(params)
	}

	return failures, nil
}
