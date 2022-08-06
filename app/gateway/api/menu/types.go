package menu

import (
	"comies/app/sdk/types"
)

type (
	Ingredient struct {
		ID           string         `json:"id"`
		ProductID    string         `json:"product_id"`
		IngredientID string         `json:"ingredient_id"`
		Quantity     types.Quantity `json:"quantity"`
		Optional     bool           `json:"optional"`
	}

	AdditionResult struct {
		ID string `json:"id"`
	}
)
