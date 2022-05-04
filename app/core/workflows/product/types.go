package product

import (
	"gomies/app/core/entities/catalog/product"
	"gomies/app/sdk/types"
)

type (
	IngredientList []product.Ingredient

	IngredientInput struct {
		Quantity     types.Quantity
		IngredientID types.ID
	}

	Reservation struct {
		ID        types.ID
		ProductID types.ID
		Quantity  types.Quantity
		Price     types.Currency
		Ignore    []types.ID
		Replace   map[types.ID]types.ID
		composite bool
	}

	ReservationResult struct {
		Price        types.Currency
		FailedChecks []FailedReservation
	}

	FailedReservation struct {
		ProductID types.ID
		Want      types.Quantity
		Got       types.Quantity
		Error     error
	}
)
