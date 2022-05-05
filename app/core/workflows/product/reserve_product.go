package product

import (
	"context"
	"gomies/app/core/entities/catalog/product"
	"gomies/app/sdk/fault"
)

func (w workflow) ReserveProduct(ctx context.Context, reservation Reservation) (Reservation, error) {
	const operation = "Workflows.Product.ReserveProduct"

	var (
		results []ItemFailed
		err     error
	)

	ingredients, err := w.products.ListIngredients(ctx, product.Key{ID: reservation.ProductID})
	if err != nil {
		return Reservation{}, fault.Wrap(err, operation)
	}

	if len(ingredients) == 0 {
		results, err = w.stocks.ReserveResources(ctx, reservation.ID, product.Ingredient{
			Quantity:     reservation.Quantity,
			IngredientID: reservation.ProductID,
		})
	} else {
		// This block of code removes from the ingredient array ignored ingredients and
		// replaces the ones set as replaced in the reservation
		var ing []product.Ingredient
		for _, ingredient := range ingredients {
			if sub, ok := reservation.Replace[ingredient.IngredientID]; ok {
				ing = append(ing, product.Ingredient{
					IngredientID: sub,
					Quantity:     ingredient.Quantity * reservation.Quantity,
				})
				continue
			}

			var shouldIgnore bool
			for _, ignore := range reservation.Ignore {
				if ignore == ingredient.IngredientID {
					shouldIgnore = true
					break
				}
			}

			if !shouldIgnore {
				ing = append(ing, product.Ingredient{
					IngredientID: ingredient.IngredientID,
					Quantity:     ingredient.Quantity * reservation.Quantity,
				})
			}
		}

		results, err = w.stocks.ReserveResources(ctx, reservation.ID, ing...)
	}

	if err != nil {
		return Reservation{}, fault.Wrap(err, operation)
	}

	reservation.Failures = results

	return reservation, nil
}
