package menu

import (
	"comies/app/core/entities/ingredient"
	"comies/app/sdk/types"
)

type (
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
		Ignore    ingredient.IgnoredList
		Replace   ingredient.ReplacedList
		Failures  []ItemFailed
	}

	ItemFailed struct {
		ProductID types.ID
		Want      types.Quantity
		Got       types.Quantity
		Error     error
	}
)
