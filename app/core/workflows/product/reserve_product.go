package product

import (
	"context"
	"gomies/app/core/entities/catalog/product"
	"gomies/app/sdk/fault"
)

func (w workflow) ReserveProduct(ctx context.Context, reservation Reservation) (ReservationResult, error) {
	const operation = "Workflows.Product.CheckProduct"

	reservation, err := w.checkSaleProperties(ctx, reservation)
	if err != nil {
		return ReservationResult{}, fault.Wrap(err, operation)
	}

	failed, err := w.reserveProductStock(ctx, reservation)
	if err != nil {
		return ReservationResult{}, fault.Wrap(err, operation)
	}

	return ReservationResult{
		Price:        reservation.Price,
		FailedChecks: failed,
	}, nil

}

func (w workflow) checkSaleProperties(ctx context.Context, reservation Reservation) (Reservation, error) {
	const operation = "Workflows.Product.approveSaleProperties"

	saleProps, err := w.products.GetProductSaleInfo(ctx, product.Key{ID: reservation.ProductID})
	if err != nil {
		return Reservation{}, fault.Wrap(err, operation)
	}

	if saleProps.SalePrice != reservation.Price {
		return Reservation{}, fault.Wrap(product.ErrInvalidSalePrice, operation)
	}

	if saleProps.MinimumSale > reservation.Quantity {
		return Reservation{}, fault.Wrap(product.ErrInvalidSaleQuantity, operation)
	}

	reservation.composite = saleProps.HasIngredients

	return reservation, nil
}

func (w workflow) reserveProductStock(ctx context.Context, reservation Reservation) ([]FailedReservation, error) {
	const operation = "Workflows.Product.reserveProductStock"

	var (
		results []FailedReservation
		err     error
	)
	if !reservation.composite {
		results, err = w.stocks.ReserveResources(ctx, reservation.ID, product.Ingredient{
			Quantity:     reservation.Quantity,
			IngredientID: reservation.ProductID,
		})
	} else {
		ingredients, err := w.products.ListIngredients(ctx, product.Key{ID: reservation.ProductID})
		if err != nil {
			return []FailedReservation{}, fault.Wrap(err, operation)
		}

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
		return []FailedReservation{}, fault.Wrap(err, operation)
	}

	return results, nil
}
