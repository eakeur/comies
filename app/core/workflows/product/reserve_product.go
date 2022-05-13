package product

import (
	"context"
	"gomies/app/core/entities/ingredient"
	"gomies/app/sdk/fault"
)

func (w workflow) ReserveProduct(ctx context.Context, r Reservation) (Reservation, error) {
	const operation = "Workflows.Product.ReserveProduct"

	var (
		failures []ItemFailed
		err      error
	)

	ingredients, err := w.ingredients.ListIngredients(ctx, r.ProductID)
	if err != nil {
		return Reservation{}, fault.Wrap(err, operation, fault.AdditionalData{
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
		return Reservation{}, fault.Wrap(err, operation)
	}

	r.Failures = failures

	return r, nil
}
