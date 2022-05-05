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

	Check struct {
		ProductID types.ID
		Quantity  types.Quantity
		Price     types.Currency
	}

	Reservation struct {
		ID        types.ID
		ProductID types.ID
		Quantity  types.Quantity
		Ignore    []types.ID
		Replace   map[types.ID]types.ID
		Failures  []ItemFailed
		composite bool
	}

	ItemFailed struct {
		ProductID types.ID
		Want      types.Quantity
		Got       types.Quantity
		Error     error
	}
)
