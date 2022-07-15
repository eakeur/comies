package menu

import (
	"comies/app/core/entities/ingredient"
	"comies/app/sdk/fault"
	"context"
)

func (w workflow) ReserveProduct(ctx context.Context, r Reservation) (Reservation, error) {

	var (
		failures []ItemFailed
		err      error
	)

	ingredients, err := w.ingredients.List(ctx, r.ProductID)
	if err != nil {
		return Reservation{}, fault.Wrap(err).Params(map[string]interface{}{
			"product_id": r.ProductID,
		})
	}

	if len(ingredients) == 0 {
		failures, err = w.stocks.ReserveResources(ctx, r.ID, ingredient.Ingredient{
			Quantity:     r.Quantity,
			IngredientID: r.ProductID,
		})
	} else {
		failures, err = w.stocks.ReserveResources(ctx, r.ID,
			ingredient.IgnoreAndReplace(
				ingredients, r.Ignore, r.Replace,
				func(i ingredient.Ingredient) ingredient.Ingredient {
					i.Quantity *= r.Quantity
					return i
				},
			)...,
		)
	}

	if err != nil {
		return Reservation{}, fault.Wrap(err)
	}

	r.Failures = failures

	return r, nil
}
